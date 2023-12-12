package handlers

import (
	"fmt"
	"strconv"

	"github.com/geraldo-labs/merge-struct"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nixpig/nixpigweb/internal/pkg/models"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
	"github.com/nixpig/nixpigweb/internal/pkg/services"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	var err error

	token := c.Locals("user")
	if token == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	claims := token.(*jwt.Token).Claims.(jwt.MapClaims)

	if !services.ValidateUserToken(token.(*jwt.Token)) {
		fmt.Println("ERROR: failed to validate user token")

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorised",
			"data":    nil,
		})
	}

	// TODO: handle the scenario where there's no "user_id" field on claims
	loggedInUserId := int(claims["user_id"].(float64))

	loggedInUser, err := queries.GetUserById(loggedInUserId)
	if err != nil || !loggedInUser.IsAdmin {
		fmt.Println("err: ", err)
		fmt.Println("isAdmin: ", loggedInUser.IsAdmin)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	var newUser models.User

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad user data provided",
			"data":    nil,
		})
	}

	validate := validator.New()

	if err := validate.Struct(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad data not validated",
			"data":    nil,
		})
	}

	_, err = queries.GetUserByUsername(newUser.Username)
	// we wan't there to be an err (i.e. no user by that name already)
	fmt.Println("err:", err)
	if err == nil {
		fmt.Println("should error")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "username already taken",
			"data":    nil,
		})
	}

	_, err = queries.GetUserByEmail(newUser.Email)
	// we wan't there to be an err (i.e. no user with that email already)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "email already taken",
			"data":    nil,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "some error on our end",
			"data":    nil,
		})
	}

	newUser.Password = string(hashedPassword)

	addedRows, err := queries.CreateUser(&newUser)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "something went wrong on our end",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("added %v users", addedRows),
		"data":    nil,
	})
}

func GetUsers(c *fiber.Ctx) error {
	users, err := queries.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to retrieve users; check your request",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("got %v users", len(users)),
		"data":    users,
	})
}

func GetUserById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request, couldn't get user",
			"data":    nil,
		})
	}

	user, err := queries.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "couldn't get user requested",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "found user",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	var updatedUser *models.User

	if err := c.BodyParser(&updatedUser); err != nil {
		fmt.Println(fmt.Errorf("failed to parse updated user\n%v", err))

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to parse id\n%v", err))

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "internal server error",
			"data":    nil,
		})
	}

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	loggedInUserId := int(claims["user_id"].(float64))

	if !services.ValidateUserToken(token) {
		fmt.Println(fmt.Errorf("unable to validate token"))

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorised",
			"data":    nil,
		})
	}

	existingUser, err := queries.GetUserById(id)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to get existing user by id\n%v", err))

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "internal server error",
			"data":    nil,
		})
	}

	if loggedInUserId != id || !existingUser.IsAdmin {
		fmt.Println(fmt.Errorf("logged in user does not match user to update or user is not an admin"))

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	patchedUser, err := mp.Struct(&existingUser, updatedUser)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to patch user struct\n%v", err))

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "internal server error",
			"data":    nil,
		})
	}

	if !patchedUser {
		fmt.Println(fmt.Errorf("no changes to patch"))

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	_, err = queries.UpdateUser(&existingUser)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to update user\n%v", err))

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "internal server error",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "successfully updated user",
		"data":    nil,
	})
}

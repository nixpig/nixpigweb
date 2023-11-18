package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nixpig/nixpigweb/internal/pkg/models"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
)

func CreateUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
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

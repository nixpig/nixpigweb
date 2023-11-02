package controllers

import (
	"fmt"
	"strconv"

	"github.com/geraldo-labs/merge-struct"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"github.com/nixpig/nixpigweb/api/database"
	"github.com/nixpig/nixpigweb/api/models"
	"github.com/nixpig/nixpigweb/api/utils"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.Connect()

	users, err := db.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "no users found",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("found %v users", len(users)),
		"data":    users,
	})
}

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "unable to parse id",
			"data":    nil,
		})
	}

	db := database.Connect()

	user, err := db.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "user with the provided id was not found",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "found user",
		"data":    user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := &models.NewUser{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to parse user from body",
			"data":    nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to validate user data",
			"data":    nil,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "unable to create user",
			"data":    nil,
		})
	}

	user.Password = string(hashedPassword)

	db := database.Connect()
	if err := db.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "unable to insert user into database",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "user created",
		"data": models.NewUser{
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to parse user id",
			"data":    nil,
		})
	}

	db := database.Connect()

	user, err := db.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "user id not found",
			"data":    nil,
		})
	}

	token := c.Locals("user").(*jwt.Token)
	isValidUserToken := utils.ValidateUserToken(token, user.Id)

	if !isValidUserToken {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "logged in user is not authorised to perform this action",
			"data":    nil,
		})
	}

	if err := db.DeleteUser(user.Id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "unable to delete user from database",
			"data":    nil,
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to parse id",
			"data":    nil,
		})
	}

	userUpdates := &models.User{}

	if err := c.BodyParser(userUpdates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to parse user updates from request body",
			"data":    nil,
		})
	}

	db := database.Connect()

	user, err := db.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "could not find user by id",
			"data":    nil,
		})
	}

	token := c.Locals("user").(*jwt.Token)
	isValidUserToken := utils.ValidateUserToken(token, user.Id)

	if !isValidUserToken {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "logged in user is not authorised to perform this action",
			"data":    nil,
		})
	}

	patched, err := mp.Struct(&user, userUpdates)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "unable to merge changes to user",
			"data":    nil,
		})
	}

	if !patched {
		return c.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"error":   false,
			"message": "no changes to make",
			"data":    user,
		})
	}

	if err := db.UpdateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "unable to update user",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "updated user",
		"data":    user,
	})
}

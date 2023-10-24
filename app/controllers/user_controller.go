package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/queries"

	"github.com/nixpig/nixpigweb/api/database"
)

func GetUsers(c *fiber.Ctx) error {
	fmt.Println("before database")
	db := database.Connect()
	userQueries := queries.UserQueries{DB: db}

	users, err := userQueries.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "no users found",
			"count":   0,
			"users":   nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"count":   len(users),
		"users":   users,
	})
}

func GetUser(c *fiber.Ctx) error {
	fmt.Println("before database")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"count":   0,
			"user":    nil,
		})
	}

	fmt.Println("before query construction")
	db := database.Connect()
	userQueries := queries.UserQueries{DB: db}

	fmt.Println("before query execution")

	user, err := userQueries.GetUser(id)
	fmt.Println("after query execution")
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "user with the provided id was not found",
			"user":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"user":    user,
	})
}

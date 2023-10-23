package controllers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/queries"

	"github.com/nixpig/nixpigweb/api/connections"
)

func GetUsers(c *fiber.Ctx) error {
	db, err := connections.Postgres()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	userQueries := queries.UserQueries{Db: db}

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
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"count":   0,
			"user":    nil,
		})
	}
	fmt.Println("id passed id:", id)

	db, err := connections.Postgres()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	userQueries := queries.UserQueries{Db: db}

	user, err := userQueries.GetUser(id)
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

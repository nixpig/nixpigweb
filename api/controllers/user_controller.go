package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/queries"

	"github.com/nixpig/nixpigweb/api/connections"
)

func GetUsers(c *fiber.Ctx) error {
	fmt.Println("connecting to database")
	db, err := connections.Postgres()
	if err != nil {
		log.Fatal(err)
		fmt.Println("error connecting to database")
		os.Exit(1)
	}
	fmt.Println("getting users...")
	userQueries := queries.UserQueries{Db: db}

	fmt.Println("querying for users")
	users, err := userQueries.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

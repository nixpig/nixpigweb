package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/database"
)

func GetStatuses(c *fiber.Ctx) error {
	db := database.Connect()

	statuses, err := db.GetStatuses()
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "got statuses",
		"data":    statuses,
	})
}

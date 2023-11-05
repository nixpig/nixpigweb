package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/database"
)

func GetCategories(c *fiber.Ctx) error {
	db := database.Connect()

	categories, err := db.GetCategories()
	fmt.Println("err", err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "found categories",
		"data":    categories,
	})
}

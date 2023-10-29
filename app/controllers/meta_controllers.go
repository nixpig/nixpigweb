package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/database"
)

func GetMeta(c *fiber.Ctx) error {
	db := database.Connect()

	meta, err := db.GetMeta()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("found %d meta items", len(meta)),
		"data":    meta,
	})
}

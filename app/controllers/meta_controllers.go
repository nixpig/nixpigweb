package controllers

import (
	"fmt"
	"strconv"

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

func GetMetaById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	db := database.Connect()

	meta, err := db.GetMetaById(id)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "couldn't get meta item",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "found meta",
		"data":    meta,
	})

}

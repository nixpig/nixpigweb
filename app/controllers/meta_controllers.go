package controllers

import (
	"fmt"
	"strconv"

	mp "github.com/geraldo-labs/merge-struct"
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/database"
	"github.com/nixpig/nixpigweb/api/models"
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

func CreateMeta(c *fiber.Ctx) error {
	meta := models.Meta{}

	if err := c.BodyParser(&meta); err != nil {
		fmt.Println("err:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "invalid meta request",
			"data":    nil,
		})
	}

	db := database.Connect()

	if err := db.CreateMeta(meta); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "something went wrong while saving meta",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "saved meta",
		"data":    nil,
	})
}

func DeleteMeta(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	db := database.Connect()

	meta, err := db.GetMetaById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	if err := db.DeleteMeta(meta.Id); err != nil {
		fmt.Println("err:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "couldn't delete meta",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "successfully deleted meta item",
		"data":    nil,
	})
}

func UpdateMeta(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   false,
			"message": "bad request",
			"data":    nil,
		})
	}

	db := database.Connect()

	meta, err := db.GetMetaById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   false,
			"message": "bad request",
			"data":    nil,
		})
	}

	metaUpdates := &models.Meta{}

	if err := c.BodyParser(metaUpdates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   false,
			"message": "bad request",
			"data":    nil,
		})
	}

	patched, err := mp.Struct(&meta, metaUpdates)
	if err != nil || !patched {
		return c.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"error":   false,
			"message": "no changes applied",
			"data":    nil,
		})
	}

	if err := db.UpdateMeta(meta); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "failed to update meta",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "successfull applied changes",
		"data":    nil,
	})
}

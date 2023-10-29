package controllers

import (
	"fmt"
	"strconv"

	mp "github.com/geraldo-labs/merge-struct"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nixpig/nixpigweb/api/database"
	"github.com/nixpig/nixpigweb/api/models"
	"github.com/nixpig/nixpigweb/api/utils"
)

func GetConfigs(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	isAdminToken := utils.ValidateAdminToken(token)

	if !isAdminToken {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "you are not authorised",

			"data": nil,
		})
	}

	db := database.Connect()

	config, err := db.GetConfigs()
	if err != nil {
		fmt.Println("err:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "no configs found",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("found %d configs", len(config)),
		"data":    config,
	})
}

func GetConfig(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	isAdminToken := utils.ValidateAdminToken(token)

	if !isAdminToken {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "you are not authorised",

			"data": nil,
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "could not parse provided id",
			"data":    nil,
		})
	}

	db := database.Connect()

	config, err := db.GetConfig(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "error fetching config",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "found config",
		"data":    config,
	})
}

func CreateConfig(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	isAdminToken := utils.ValidateAdminToken(token)

	if !isAdminToken {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "you are not authorised",

			"data": nil,
		})
	}

	config := models.NewConfig{}

	if err := c.BodyParser(&config); err != nil {
		fmt.Println("err", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "could not parse provided config",
			"data":    nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		fmt.Println("failed validation:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to validate config",
			"data":    nil,
		})
	}

	db := database.Connect()

	if err := db.CreateConfig(&config); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "could no create config",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "successfully created config",
		"data":    nil,
	})
}

func DeleteConfig(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	isAdminToken := utils.ValidateAdminToken(token)

	if !isAdminToken {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "you are not authorised",

			"data": nil,
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "you provided a bad config id to delete",
			"data":    nil,
		})
	}

	db := database.Connect()

	if err := db.DeleteConfig(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "encountered an error while deleting config",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("successfully deleted config with id: %v", id),
	})
}

func UpdateConfig(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	isAdminToken := utils.ValidateAdminToken(token)

	if !isAdminToken {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "you are not authorised",

			"data": nil,
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": fmt.Sprintf("sorry, could not parse provided id: %v", id),
			"data":    nil,
		})
	}

	configUpdates := models.Config{}

	if err := c.BodyParser(&configUpdates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "failed trying to parse config updates",
			"data":    nil,
		})
	}

	db := database.Connect()

	config, err := db.GetConfig(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad id",
			"data":    nil,
		})
	}

	patched, err := mp.Struct(&config, configUpdates)
	if err != nil || !patched {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "could not merge in proposed changes",
			"data":    nil,
		})
	}

	if err := db.UpdateConfig(&config); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "could not save proposed changes",
			"data":    nil,
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "successfully update config",
		"data":    nil,
	})
}

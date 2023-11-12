package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/database"
	"github.com/nixpig/nixpigweb/internal/pkg/models"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"

	"github.com/go-playground/validator/v10"

	"github.com/mozillazg/go-slugify"
)

func GetContent(c *fiber.Ctx) error {
	contentQueries := queries.Content{DB: database.Connection()}

	content, err := contentQueries.GetContent()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"content": nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("found %v records", len(content)),
		"content": content,
	})
}

func CreateContent(c *fiber.Ctx) error {
	content := &models.Content{}

	if err := c.BodyParser(content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to parse body content",
			"content": nil,
		})
	}

	content.Slug = slugify.Slugify(content.Title)

	validate := validator.New()

	if err := validate.Struct(content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "could not validate data",
			"content": nil,
		})
	}

	contentQueries := queries.Content{DB: database.Connection()}

	res, err := contentQueries.CreateContent(content)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "it's not you, it's me",
			"content": nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "successfully added",
		"content": res,
	})
}

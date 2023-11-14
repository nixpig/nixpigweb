package handlers

import (
	"fmt"
	"strconv"
	"time"

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

func GetContentById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad content id provided in request",
			"content": nil,
		})
	}

	contentQueries := queries.Content{DB: database.Connection()}

	content, err := contentQueries.GetContentById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": fmt.Sprintf("no content found with id provided: %v", id),
			"content": nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("found content for id: %v", id),
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

	rowsAffected, err := contentQueries.CreateContent(content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "it's not you, it's me",
			"content": nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("%v records added", rowsAffected),
		"content": nil,
	})
}

func DeleteContentById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad content id provided",
			"content": nil,
		})
	}

	contentQueries := queries.Content{DB: database.Connection()}

	rowsAffected, err := contentQueries.DeleteContentById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "we messed something up",
			"content": nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("%v records deleted", rowsAffected),
		"content": nil,
	})
}

func UpdateContent(c *fiber.Ctx) error {
	var content models.Content

	if err := c.BodyParser(&content); err != nil {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "bad request",
				"content": nil,
			})
		}
	}

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"content": nil,
		})
	}

	if id != content.Id {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"content": nil,
		})
	}

	validate := validator.New()

	slug := slugify.Slugify(content.Title)
	content.Slug = slug

	content.UpdatedAt = time.Now()

	if err := validate.Struct(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"content": nil,
		})
	}

	contentQueries := queries.Content{DB: database.Connection()}

	rowsAffected, err := contentQueries.UpdateContent(&content)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "we messed up",
			"content": nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("%v records updated", rowsAffected),
		"content": nil,
	})
}

package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/models"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
	"github.com/nixpig/nixpigweb/internal/pkg/services"

	"github.com/go-playground/validator/v10"

	"github.com/mozillazg/go-slugify"
)

func GetContent(c *fiber.Ctx) error {
	content, err := queries.GetContent()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("found %v records", len(content)),
		"data":    content,
	})
}

func GetContentById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad content id provided in request",
			"data":    nil,
		})
	}

	content, err := queries.GetContentById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": fmt.Sprintf("no content found with id provided: %v", id),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("found content for id: %v", id),
		"data":    content,
	})
}

func CreateContent(c *fiber.Ctx) error {
	content := &models.Content{}

	if err := c.BodyParser(content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "unable to parse body content",
			"data":    nil,
		})
	}

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	if !services.ValidateUserToken(token, userId) {
		fmt.Println("ERROR: failed to validate user token")

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorised",
			"data":    nil,
		})
	}

	user, err := queries.GetUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "user does not exist",
			"data":    nil,
		})
	}

	content.UserId = user.Id
	content.Slug = slugify.Slugify(content.Title)

	validate := validator.New()

	if err := validate.Struct(content); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "could not validate data",
			"data":    nil,
		})
	}

	rowsAffected, err := queries.CreateContent(content)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "it's not you, it's me",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("%v records added", rowsAffected),
		"data":    nil,
	})
}

func DeleteContentById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad content id provided",
			"data":    nil,
		})
	}

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	loggedInUserId := int(claims["user_id"].(float64))

	loggedInUser, err := queries.GetUserById(loggedInUserId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	content, err := queries.GetContentById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	if loggedInUser.Id != content.UserId && !loggedInUser.IsAdmin {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	rowsAffected, err := queries.DeleteContentById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "we messed something up",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("%v records deleted", rowsAffected),
		"data":    nil,
	})
}

func UpdateContent(c *fiber.Ctx) error {
	var updatedContent models.Content

	if err := c.BodyParser(&updatedContent); err != nil {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "bad request",
				"data":    nil,
			})
		}
	}

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	loggedInUserId := int(claims["user_id"].(float64))

	existingContent, err := queries.GetContentById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	loggedInUser, err := queries.GetUserById(loggedInUserId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	if (loggedInUserId != updatedContent.UserId || loggedInUserId != existingContent.UserId) && !loggedInUser.IsAdmin {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	if id != updatedContent.Id {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	validate := validator.New()

	slug := slugify.Slugify(updatedContent.Title)
	updatedContent.Slug = slug

	updatedContent.UpdatedAt = time.Now()

	if err := validate.Struct(&updatedContent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "bad request",
			"data":    nil,
		})
	}

	rowsAffected, err := queries.UpdateContent(&updatedContent)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "we messed up",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("%v records updated", rowsAffected),
		"data":    nil,
	})
}

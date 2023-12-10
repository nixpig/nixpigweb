package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
)

func Protected() fiber.Handler {
	return jwtware.New(
		jwtware.Config{
			SigningKey:   jwtware.SigningKey{Key: []byte(config.Get("SECRET"))},
			ErrorHandler: jwtErrorHandler,
		},
	)
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	errorMessage := err.Error()

	if errorMessage == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": errorMessage,
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   true,
		"message": "not authorised",
		"data":    nil,
	})
}

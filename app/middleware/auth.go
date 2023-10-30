package middleware

import (
	"fmt"

	"github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nixpig/nixpigweb/api/config"
	"github.com/nixpig/nixpigweb/api/utils"
)

func Admin(c *fiber.Ctx) error {
	fmt.Println("in admin middleware")
	token := c.Locals("user").(*jwt.Token)
	if utils.ValidateAdminToken(token) {
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   true,
		"message": "you need to be admin",
		"data":    nil,
	})

}

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.Get("SECRET"))},
		ErrorHandler: jwtErrorHandler,
	})
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "invalid jwt provided",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   true,
		"message": "not authorised",
		"data":    nil,
	})
}

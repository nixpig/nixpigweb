package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/api/handlers"
)

func RegisterAuthRoutes(api fiber.Router) fiber.Router {
	auth := api.Group("/auth")

	auth.Post("/login", handlers.Login)

	return auth
}

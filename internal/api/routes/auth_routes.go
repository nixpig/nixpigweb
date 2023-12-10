package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/api/handlers"
	"github.com/nixpig/nixpigweb/internal/api/middleware"
)

func RegisterAuthRoutes(api fiber.Router) fiber.Router {
	auth := api.Group("/auth")

	auth.Post("/login", handlers.Login)
	auth.Post("/logout", middleware.Protected(), handlers.Logout)
	auth.Post("/update-password", middleware.Protected(), handlers.ChangePassword)

	return auth
}

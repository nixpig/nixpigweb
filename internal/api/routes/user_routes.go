package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/api/handlers"
	"github.com/nixpig/nixpigweb/internal/api/middleware"
)

func RegisterUserRoutes(api fiber.Router) fiber.Router {
	user := api.Group("/user")

	user.Post("/", middleware.Protected(), handlers.CreateUser)
	user.Get("/", handlers.GetUsers)
	user.Get("/:id", handlers.GetUserById)

	return user
}

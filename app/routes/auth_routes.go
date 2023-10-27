package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/controllers"
)

func SetupAuthRoutes(api fiber.Router) {
	auth := api.Group("/auth")

	auth.Post("/", controllers.Login)
}

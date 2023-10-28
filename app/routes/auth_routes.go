package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/controllers"
)

func SetupAuthRoutes(api fiber.Router) fiber.Router {
	auth := api.Group("/login")

	auth.Post("/", controllers.Login)

	return auth
}

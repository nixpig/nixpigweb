package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/controllers"
)

func SetupStatusRoutes(api fiber.Router) fiber.Router {
	status := api.Group("/status")

	status.Get("/", controllers.GetStatuses)

	return status
}

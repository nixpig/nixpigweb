package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/controllers"
	"github.com/nixpig/nixpigweb/api/middleware"
)

func SetupConfigRoutes(api fiber.Router) fiber.Router {
	config := api.Group("/config")

	config.Get("/", middleware.Protected(), controllers.GetConfigs)
	config.Get("/:id", middleware.Protected(), controllers.GetConfig)
	config.Post("/", middleware.Protected(), controllers.CreateConfig)
	config.Patch("/:id", middleware.Protected(), controllers.UpdateConfig)
	config.Delete("/:id", middleware.Protected(), controllers.DeleteConfig)

	return config
}

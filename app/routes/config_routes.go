package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/controllers"
)

func SetupConfigRoutes(api fiber.Router) fiber.Router {
	config := api.Group("/config")

	config.Get("/", controllers.GetConfigs)
	config.Get("/id", controllers.GetConfig)
	config.Post("/", controllers.CreateConfig)
	config.Patch("/:id", controllers.UpdateConfig)
	config.Delete("/:id", controllers.DeleteConfig)

	return config
}

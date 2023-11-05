package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/controllers"
)

func SetupCategoryRoutes(api fiber.Router) fiber.Router {
	category := api.Group("/category")

	category.Get("/", controllers.GetCategories)

	return category
}

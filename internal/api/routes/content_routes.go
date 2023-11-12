package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/api/handlers"
)

func RegisterContentRoutes(api fiber.Router) fiber.Router {
	content := api.Group("/content")

	content.Get("/", handlers.GetContent)
	content.Get("/:id", handlers.GetContentById)
	content.Post("/", handlers.CreateContent)

	return content
}

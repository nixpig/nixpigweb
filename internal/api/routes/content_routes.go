package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/api/handlers"
	"github.com/nixpig/nixpigweb/internal/api/middleware"
)

func RegisterContentRoutes(api fiber.Router) fiber.Router {
	content := api.Group("/content")

	content.Get("/", handlers.GetContent)
	content.Post("/", middleware.Protected(), handlers.CreateContent)
	content.Patch("/:id", middleware.Protected(), handlers.UpdateContent)
	content.Get("/:id", handlers.GetContentById)
	content.Delete("/:id", middleware.Protected(), handlers.DeleteContentById)

	return content
}

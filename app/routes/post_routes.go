package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/controllers"
	"github.com/nixpig/nixpigweb/api/middleware"
)

func SetupPostRoutes(api fiber.Router) {
	post := api.Group("/post")

	post.Get("/", controllers.GetPosts)
	post.Get("/:id", controllers.GetPost)
	post.Delete("/:id", middleware.Protected(), controllers.DeletePost)
	post.Post("/", middleware.Protected(), controllers.CreatePost)
	post.Patch("/:id", middleware.Protected(), controllers.UpdatePost)
}

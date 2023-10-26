package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/controllers"
)

func SetupPostRoutes(api fiber.Router) {
	post := api.Group("/post")

	post.Get("/", controllers.GetPosts)
	post.Get("/:id", controllers.GetPost)
	post.Delete("/:id", controllers.DeletePost)
	post.Post("/", controllers.CreatePost)

}

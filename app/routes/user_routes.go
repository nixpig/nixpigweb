package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/controllers"
	"github.com/nixpig/nixpigweb/api/middleware"
)

func SetupUserRoutes(api fiber.Router) fiber.Router {
	user := api.Group("/user")

	user.Get("/", controllers.GetUsers)
	user.Get("/:id", controllers.GetUser)
	user.Post("/register", controllers.CreateUser)
	user.Delete("/:id", middleware.Protected(), controllers.DeleteUser)
	user.Patch("/:id", middleware.Protected(), controllers.UpdateUser)

	return user
}

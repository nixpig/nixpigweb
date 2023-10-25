package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/controllers"
)

func SetupUserRoutes(api fiber.Router) {
	user := api.Group("/user")

	user.Get("/", controllers.GetUsers)
	user.Get("/:id", controllers.GetUser)
	user.Post("/", controllers.CreateUser)
	user.Delete("/:id", controllers.DeleteUser)
	user.Patch("/:id", controllers.UpdateUser)
}

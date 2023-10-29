package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/controllers"
	"github.com/nixpig/nixpigweb/api/middleware"
)

func SetupMetaRoutes(api fiber.Router) fiber.Router {

	meta := api.Group("/meta")

	meta.Get("/", middleware.Protected(), controllers.GetMeta)
	meta.Get("/:id", middleware.Protected(), controllers.GetMetaById)

	return meta
}

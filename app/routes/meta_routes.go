package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/controllers"
	"github.com/nixpig/nixpigweb/api/middleware"
)

func SetupMetaRoutes(api fiber.Router) fiber.Router {

	meta := api.Group(
		"/meta",
		middleware.Protected(),
		middleware.Admin,
	)

	meta.Get("/", controllers.GetMeta)
	meta.Get("/:id", controllers.GetMetaById)
	meta.Post("/", controllers.CreateMeta)
	meta.Delete("/:id", controllers.DeleteMeta)

	return meta
}

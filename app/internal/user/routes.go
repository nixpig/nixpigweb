package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/internal/database"
)

func SetupUserRoutes(api fiber.Router) {
	user := api.Group("/user")
	db := database.Connect()
	queries := UserQueries{DB: db}

	user.Get("/", GetAllHandler(queries))
	user.Get("/:id", GetOneHandler(queries))
	user.Post("/", CreateHandler(queries))
	user.Delete("/:id", DeleteHandler(queries))
	user.Patch("/:id", UpdateHandler(queries))
}

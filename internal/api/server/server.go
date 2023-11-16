package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nixpig/nixpigweb/internal/api/routes"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
)

func Start() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	api := app.Group("/api")

	routes.RegisterContentRoutes(api)
	routes.RegisterUserRoutes(api)

	api.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "not found",
			"data":    nil,
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Get("API_PORT"))))
}

package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/nixpig/nixpigweb/internal/api/routes"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
)

func Start(contextPath string, port string) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return config.Get("APP_ENV") == "development"
		},
	}))

	app.Use(helmet.New())
	app.Use(logger.New())

	api := app.Group(fmt.Sprintf("/%s", contextPath))

	routes.RegisterContentRoutes(api)
	routes.RegisterUserRoutes(api)
	routes.RegisterAuthRoutes(api)

	api.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "not found",
			"data":    nil,
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}

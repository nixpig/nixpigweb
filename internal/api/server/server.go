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

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Get("API_PORT"))))
}

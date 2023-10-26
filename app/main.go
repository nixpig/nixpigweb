package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nixpig/nixpigweb/api/config"
	"github.com/nixpig/nixpigweb/api/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	api := app.Group("/api")

	routes.SetupUserRoutes(api)
	routes.SetupPostRoutes(api)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Get("API_PORT"))))
}
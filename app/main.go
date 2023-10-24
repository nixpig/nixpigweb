package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nixpig/nixpigweb/api/routes"
)

func main() {
	app := fiber.New()

	api := app.Group("/api", logger.New())

	routes.SetupUserRoutes(api)

	log.Fatal(app.Listen(":3000"))
}

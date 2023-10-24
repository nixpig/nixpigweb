package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/api/routes"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	routes.SetupUserRoutes(api)

	log.Fatal(app.Listen(":3000"))
}

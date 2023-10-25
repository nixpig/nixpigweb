package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nixpig/nixpigweb/api/internal/user"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	api := app.Group("/api")

	user.SetupUserRoutes(api)

	log.Fatal(app.Listen(":3000"))
}

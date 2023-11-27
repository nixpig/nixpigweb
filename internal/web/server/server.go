package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	"github.com/nixpig/nixpigweb/internal/web/handlers"
)

func Start(contextPath string, port string) {
	engine := html.New("./internal/web/templates/", ".html")

	// TODO: dynamically set development settings below
	engine.Reload(true)
	engine.Debug(true)

	engine.AddFunc("greet", func(name string) string {
		return fmt.Sprintf("Hello, %s!", name)
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(logger.New())

	web := app.Group(fmt.Sprintf("/%s", contextPath))

	web.Get("/", handlers.IndexHandler)

	web.Use(handlers.NotFoundHandler)

	// TODO: add 404

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))

}

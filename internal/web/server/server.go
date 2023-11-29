package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/web/handlers"
	"log"
)

func Start(contextPath string, port string) {
	engine := html.New("./web/templates/", ".tmpl")

	env := config.Get("APP_ENV")

	if env == "development" {
		engine.Reload(true)
		engine.Debug(true)
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./web/static")

	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(compress.New())

	web := app.Group(fmt.Sprintf("/%s", contextPath))

	web.Get("/", handlers.IndexHandler)
	web.Get("/blog", handlers.BlogHander)
	web.Get("/:slug", handlers.ContentHandler)

	web.Use(handlers.NotFoundHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}

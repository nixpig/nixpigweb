package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nixpig/nixpigweb/internal/web/templates"
	"github.com/nixpig/nixpigweb/internal/web/utils"
)

func Start(contextPath string, port string) {
	app := fiber.New()

	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(logger.New())

	web := app.Group(fmt.Sprintf("/%s", contextPath))

	web.Get("/", utils.WrapComponent(templates.Hello("pig")))

	// TODO: add 404

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))

}

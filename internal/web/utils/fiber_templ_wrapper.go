package utils

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func WrapComponent(c templ.Component) fiber.Handler {
	return adaptor.HTTPHandler(templ.Handler(c))
}

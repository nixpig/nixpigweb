package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
)

func NotFoundHandler(c *fiber.Ctx) error {
	contextPath := config.Get("WEB_CONTEXT")
	siteName := config.Get("SITE_NAME")

	return c.Render("404", fiber.Map{
		"SiteName":    siteName,
		"ContextPath": contextPath,
	})

}

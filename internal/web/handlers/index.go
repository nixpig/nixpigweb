package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
)

func IndexHandler(c *fiber.Ctx) error {
	content, err := queries.GetContent()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("500", fiber.Map{

			"SiteName":    "nixpig.dev",
			"ContextPath": config.Get("WEB_CONTEXT"),
		})
	}

	return c.Render("index", fiber.Map{
		"Title":   "Hello, world!",
		"Content": content,
	})
}

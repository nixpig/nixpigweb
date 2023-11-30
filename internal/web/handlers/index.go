package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
)

func IndexHandler(c *fiber.Ctx) error {
	pages, err := queries.GetContentByType("page")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("500", fiber.Map{
			"PageTitle":   "500 - Internal Server Error",
			"SiteName":    config.Get("SITE_NAME"),
			"ContextPath": config.Get("WEB_CONTEXT"),
		})
	}

	return c.Render("index", fiber.Map{
		"PageTitle":   "Index Page Title",
		"Pages":       pages,
		"ContextPath": config.Get("WEB_CONTEXT"),
		"SiteName":    config.Get("SITE_NAME"),
	})
}

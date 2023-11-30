package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
)

func NotFoundHandler(c *fiber.Ctx) error {
	contextPath := config.Get("WEB_CONTEXT")
	siteName := config.Get("SITE_NAME")
	pageTitle := "404 - Not found"

	pages, err := queries.GetContentByType("page")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("500", fiber.Map{
			"PageTitle":   pageTitle,
			"SiteName":    config.Get("SITE_NAME"),
			"ContextPath": config.Get("WEB_CONTEXT"),
		})
	}

	return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{
		"PageTitle":   pageTitle,
		"SiteName":    siteName,
		"ContextPath": contextPath,
		"Pages":       pages,
	})

}

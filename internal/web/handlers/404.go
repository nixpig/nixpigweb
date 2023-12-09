package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
)

func NotFoundHandler(c *fiber.Ctx) error {
	siteName := config.Get("SITE_NAME")
	pageTitle := "404 - Not found"

	pages, err := queries.GetContentByType("page")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("500", fiber.Map{
			"PageTitle": "500 - Internal Server Error",
			"SiteName":  config.Get("SITE_NAME"),
		})
	}

	return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{
		"PageTitle": pageTitle,
		"SiteName":  siteName,
		"Pages":     pages,
	})

}

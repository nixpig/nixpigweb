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
			"PageTitle": "500 - Internal Server Error",
			"SiteName":  config.Get("SITE_NAME"),
		})
	}

	posts, err := queries.GetContentByType("post")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("500", fiber.Map{
			"PageTitle": "500 - Internal Server Error",
			"SiteName":  config.Get("SITE_NAME"),
		})
	}

	return c.Render("index", fiber.Map{
		"Pages":        pages,
		"Posts":        posts,
		"SiteName":     config.Get("SITE_NAME"),
		"PageTitle":    "Welcome to nixpig.dev",
		"PageSubtitle": "Diary of a front-ender learning Rust, Go and other back-end tech...",
	})
}

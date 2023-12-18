package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
)

func BlogHander(c *fiber.Ctx) error {
	posts, err := queries.GetContentByType("post")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("500", fiber.Map{
			"PageTitle": "500 - Internal Server Error",
			"SiteName":  config.Get("SITE_NAME"),
		})
	}

	pages, err := queries.GetContentByType("page")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("500", fiber.Map{
			"PageTitle": "500 - Internal Server Error",
			"SiteName":  config.Get("SITE_NAME"),
		})
	}

	return c.Render("blog", fiber.Map{
		"SiteName":  config.Get("SITE_NAME"),
		"Pages":     pages,
		"PageTitle": "Blog",
		"Posts":     posts,
		"Body":      "A collection of the blog posts I've written.",
	})
}

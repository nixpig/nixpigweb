package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
)

func ContentHandler(c *fiber.Ctx) error {
	slug := c.Params("slug")

	content, err := queries.GetContentBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("500", fiber.Map{
			"PageTitle": "500 - Internal Server Error",
			"SiteName":  "nixpig.dev",
		})
	}

	pages, err := queries.GetContentByType("page")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("500", fiber.Map{
			"PageTitle": "500 - Internal Server Error",
			"SiteName":  config.Get("SITE_NAME"),
		})
	}

	return c.Render("content", fiber.Map{
		"SiteName":     "nixpig.dev",
		"Pages":        pages,
		"Id":           content.Id,
		"PageTitle":    content.Title,
		"PageSubtitle": content.Subtitle,
		"Slug":         content.Slug,
		"Body":         content.Body,
		"CreatedAt":    content.CreatedAt,
		"UpdatedAt":    content.UpdatedAt,
		"Type":         content.Type,
		"UserId":       content.UserId,
	})

}

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
		return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{
			"SiteName":    "nixpig.dev",
			"ContextPath": config.Get("WEB_CONTEXT"),
		})
	}

	return c.Render("content", fiber.Map{
		"Id":        content.Id,
		"Title":     content.Title,
		"Subtitle":  content.Subtitle,
		"Slug":      content.Slug,
		"Body":      content.Body,
		"CreatedAt": content.CreatedAt,
		"UpdatedAt": content.UpdatedAt,
		"Type":      content.Type,
		"UserId":    content.UserId,
	})

}

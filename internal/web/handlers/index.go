package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nixpig/nixpigweb/internal/pkg/config"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
)

func IndexHandler(c *fiber.Ctx) error {
	posts, err := queries.GetContentByType("post")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("500", fiber.Map{

			"SiteName":    "nixpig.dev",
			"ContextPath": config.Get("WEB_CONTEXT"),
		})
	}

	pages, err := queries.GetContentByType("page")
	if err != nil {
		return c.Status(fiber.StatusNotFound).Render("500", fiber.Map{
			"SiteName":    "nixpig.dev",
			"ContextPath": config.Get("WEB_CONTEXT"),
		})
	}

	return c.Render("index", fiber.Map{
		"Title": "Hello, world!",
		"Posts": posts,
		"Pages": pages,
	})
}

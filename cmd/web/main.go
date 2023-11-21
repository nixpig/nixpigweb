package main

import "fmt"
import "log"
import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/template/html/v2"

func main() {
	engine := html.New("./internal/web/templates/", ".html")

	// DEVELOPMENT
	engine.Reload(true)
	engine.Debug(true)

	engine.AddFunc("greet", func(name string) string {
		return fmt.Sprintf("Hello, %s!", name)
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, world!",
		})
	})

	log.Fatal(app.Listen(":3000"))

}

package public

import "github.com/gofiber/fiber/v2"

// GetDocs renders the public API documentation page.
func GetDocs(c *fiber.Ctx) error {
	return c.Render("docs", fiber.Map{
		"title": "FilmyFly API Documentation",
	})
}

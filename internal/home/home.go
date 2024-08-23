package home

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"AppName": os.Getenv("APP_NAME"),
		"Title":   "Hello, World!",
	}, "layouts/main")
}

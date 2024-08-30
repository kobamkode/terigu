package admin

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func DashboardPage(c *fiber.Ctx) error {
	return c.Render("dashboard", fiber.Map{
		"AppName": os.Getenv("APP_NAME"),
		"Title":   "Hello, Dashboard!",
	}, "layouts/main")
}

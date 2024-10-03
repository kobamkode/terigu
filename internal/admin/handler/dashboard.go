package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kobamkode/terigu/config"
)

func DashboardPage(c *fiber.Ctx) error {
	return c.Render("admin/dashboard", fiber.Map{
		"AppName": config.Get("APP_NAME"),
		"Title":   "Hello, Dashboard!",
	}, "layouts/admin")
}

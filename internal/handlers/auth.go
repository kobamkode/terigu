package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func LoginPage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"AppName": os.Getenv("APP_NAME"),
		"Title":   "Login",
	}, "layouts/auth")
}

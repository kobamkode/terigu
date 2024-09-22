package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func LoginPage(c *fiber.Ctx) error {
	return c.Render("auth", fiber.Map{
		"AppName": os.Getenv("APP_NAME"),
		"Title":   "Login",
	}, "layouts/auth")
}

func Login(c *fiber.Ctx) error {
	usr := c.FormValue("username")
	pwd := c.FormValue("password")

	if usr != "test" && pwd != "123456" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return nil
}

func Logout(c *fiber.Ctx) error {
	return nil
}

func Register(c *fiber.Ctx) error {
	return nil
}

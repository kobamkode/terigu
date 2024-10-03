package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kobamkode/terigu/config"
)

func LoginPage(c *fiber.Ctx) error {
	return c.Render("web/login", fiber.Map{
		"AppName": config.Get("APP_NAME"),
		"Title":   config.Get("APP_NAME"),
	}, "web/layout")
}

func RegisterPage(c *fiber.Ctx) error {
	return c.Render("web/register", fiber.Map{
		"AppName": config.Get("APP_NAME"),
		"Title":   config.Get("APP_NAME"),
	}, "web/layout")
}

package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kobamkode/terigu/config"
)

func HomePage(c *fiber.Ctx) error {
	return c.Render("web/home", fiber.Map{
		"AppName": config.Get("APP_NAME"),
		"Title":   config.Get("APP_NAME"),
	}, "web/layout")
}

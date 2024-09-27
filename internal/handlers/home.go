package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kobamkode/terigu/config"
)

func HomePage(c *fiber.Ctx) error {
	return c.Render("web-pages/home", fiber.Map{
		"AppName": config.Get("APP_NAME"),
		"Title":   config.Get("APP_NAME"),
	}, "layouts/web")
}

func JoinPage(c *fiber.Ctx) error {
	return c.Render("web-pages/join", fiber.Map{
		"AppName": config.Get("APP_NAME"),
		"Title":   config.Get("APP_NAME"),
	}, "layouts/web")
}

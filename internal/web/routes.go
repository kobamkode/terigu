package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kobamkode/terigu/internal/web/handler"
)

func Route(app *fiber.App) {
	app.Get("/", handler.HomePage)
	app.Get("/login", handler.LoginPage)
	app.Get("/register", handler.RegisterPage)
}

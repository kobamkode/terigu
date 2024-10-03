package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kobamkode/terigu/internal/api/handler"
)

func Route(app *fiber.App) {
	a := app.Group("/api")
	a.Get("/ping", handler.Ping)
}

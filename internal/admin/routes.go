package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kobamkode/terigu/internal/admin/handler"
)

func Route(app *fiber.App) {
	r := app.Group("/admin")
	r.Get("/", handler.DashboardPage)
}

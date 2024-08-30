package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kobamkode/terigu/internal/handlers/admin"
	"github.com/kobamkode/terigu/internal/handlers/api"
	"github.com/kobamkode/terigu/internal/handlers/web"
)

type handler struct {
	app  *fiber.App
	pool *pgxpool.Pool
}

func NewHandler(app *fiber.App, pool *pgxpool.Pool) *handler {
	return &handler{app, pool}
}

func (h *handler) Web() {
	h.app.Get("/", web.HomePage)
}

func (h *handler) API() {
	r := h.app.Group("/api")
	r.Get("/ping", api.Ping)

}

func (h *handler) Admin() {
	r := h.app.Group("/admin")
	r.Get("/", admin.DashboardPage)
	r.Get("/login", admin.LoginPage)
	// r.Get("/register", admin.RegisterPage)
}

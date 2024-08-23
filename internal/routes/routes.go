package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kobamkode/terigu/internal/home"
)

type handler struct {
	app  *fiber.App
	pool *pgxpool.Pool
}

func NewHandler(app *fiber.App, pool *pgxpool.Pool) *handler {
	return &handler{app, pool}
}

func (h *handler) Web() {
	h.app.Get("/", home.HomePage)
}

func (h *handler) API() {
	_ = h.app.Group("/api")

}

func (h *handler) Admin() {
	_ = h.app.Group("/admin")
}

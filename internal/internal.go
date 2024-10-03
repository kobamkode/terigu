package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/gofiber/storage/redis/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kobamkode/terigu/config"
	"github.com/kobamkode/terigu/internal/admin"
	"github.com/kobamkode/terigu/internal/api"
	"github.com/kobamkode/terigu/internal/web"
)

func Setup(app *fiber.App, db *pgxpool.Pool, mem *redis.Storage) {
	// Assets
	app.Static("/assets", "./views/assets")
	app.Static("/admin/assets", "./views/assets")

	// Middlewares
	// s := session.New(config.Session(store))
	// app.Use(csrf.New(config.Csrf(s)))
	app.Use(favicon.New())
	app.Use(logger.New(config.Logger()))
	app.Use(requestid.New())
	if config.Get("APP_ENV") != "local" {
		app.Use(limiter.New())

	}

	web.Route(app)
	admin.Route(app)
	api.Route(app)

}

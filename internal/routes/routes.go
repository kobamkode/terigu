package routes

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	// "github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kobamkode/terigu/config"
	"github.com/kobamkode/terigu/internal/handlers"
)

func Setup(app *fiber.App, db *pgxpool.Pool, store *redis.Storage) {
	// Assets
	app.Static("/assets", "./views/assets")
	app.Static("/admin/assets", "./views/assets")

	// Middlewares
	// s := session.New(config.Session(store))
	// app.Use(csrf.New(config.Csrf(s)))
	app.Use(logger.New(config.Logger()))
	app.Use(requestid.New())
	if config.Get("APP_ENV") != "local" {
		app.Use(limiter.New())
	}

	// Web Routes
	app.Get("/", handlers.HomePage)

	// API Routes
	api := app.Group("/api")
	api.Get("/ping", handlers.Ping)
	api.Post("/login", handlers.Login)
	api.Post("/logout", handlers.Logout)
	api.Post("/register", handlers.Register)

	// Admin Routes
	admin := app.Group("/admin")
	admin.Get("/", handlers.DashboardPage)
	admin.Get("/login", handlers.LoginPage)
}

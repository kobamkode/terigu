package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/kobamkode/terigu/config"
	"github.com/kobamkode/terigu/database"
	"github.com/kobamkode/terigu/internal/routes"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	db := database.PostgresConn()
	defer db.Close()

	store := database.RedisConn()

	routes.Setup(app, db, store)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Get("APP_PORT"))))
}

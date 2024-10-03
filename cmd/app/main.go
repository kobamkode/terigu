package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/kobamkode/terigu/config"
	"github.com/kobamkode/terigu/database"
	"github.com/kobamkode/terigu/internal"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:                 engine,
		DisableStartupMessage: true,
	})

	db := database.PostgresConn()
	defer db.Close()

	mem := database.RedisConn()

	internal.Setup(app, db, mem)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Get("APP_PORT"))))
}

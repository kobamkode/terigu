package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kobamkode/terigu/config"
)

var migrateType string

func init() {
	flag.StringVar(&migrateType, "run", "up", "run the migration up|down")
}

func main() {
	flag.Parse()

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_NAME"),
		config.Get("DB_SSL"),
	)

	m, err := migrate.New("file://gen/migrations", dbUrl)
	if err != nil {
		log.Fatal("err: " + err.Error())
	}

	if migrateType == "up" {
		if err := m.Up(); err != nil {
			log.Fatal(err.Error())
		}
	}

	if migrateType == "down" {
		if err := m.Down(); err != nil {
			log.Fatal(err.Error())
		}
	}
}

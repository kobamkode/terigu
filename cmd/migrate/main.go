package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kobamkode/terigu/internal/env"
)

var migrateType string

func init() {
	flag.StringVar(&migrateType, "run", "up", "run the migration up|down")
}

func main() {
	flag.Parse()
	env.Load()

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
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

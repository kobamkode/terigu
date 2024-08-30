package main

import (
	"github.com/kobamkode/terigu/internal/database"
	"github.com/kobamkode/terigu/internal/env"
	"github.com/kobamkode/terigu/internal/server"
)

func init() {
	env.Load()
}

func main() {
	db := database.NewDBPool()
	defer db.Close()

	server.Run(db)
}

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
	pool := database.NewDBPool()
	defer pool.Close()

	server.Run(pool)
}

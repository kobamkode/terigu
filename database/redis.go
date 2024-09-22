package database

import (
	"github.com/gofiber/storage/redis/v3"
	"github.com/kobamkode/terigu/config"
)

func RedisConn() *redis.Storage {
	return redis.New(redis.Config{
		Host:     config.Get("REDIS_HOST"),
		Port:     config.GetInt("REDIS_PORT"),
		Password: config.Get("REDIS_PASSWORD"),
		Database: 0,
		Reset:    false,
	})
}

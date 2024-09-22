package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/redis/v3"
	"github.com/joho/godotenv"
)

func Get(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetInt(key string) int {
	g := Get(key)
	i, err := strconv.Atoi(g)
	if err != nil {
		return -1
	}
	return i
}

func Session(storage *redis.Storage) session.Config {
	return session.Config{
		Storage: storage,
	}
}

func Csrf(session *session.Store) csrf.Config {
	const HeaderName = "X-Csrf-Token"
	return csrf.Config{
		KeyLookup:         "header:" + HeaderName,
		CookieName:        "__Host-csrf_",
		CookieSameSite:    "Lax",
		CookieSecure:      true,
		CookieSessionOnly: true,
		CookieHTTPOnly:    true,
		Expiration:        1 * time.Hour,
		KeyGenerator:      utils.UUIDv4,
		Extractor:         csrf.CsrfFromHeader(HeaderName),
		Session:           session,
		SessionKey:        fmt.Sprintf("%s.csrf.token", Get("APP_NAME")),
		HandlerContextKey: fmt.Sprintf("%s.csrf.handler", Get("APP_NAME")),
	}
}

func Logger() logger.Config {
	return logger.Config{
		Format: "${time} ${locals:requestid} ${latency} ${status} - ${method} ${path}\n",
	}
}

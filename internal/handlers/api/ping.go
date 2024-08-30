package api

import "github.com/gofiber/fiber/v2"

func Ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "PONG",
	})
}

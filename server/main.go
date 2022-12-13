package main

import (
	"github.com/beezy-dev/echo-ca/server/pkg/types"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	data := types.Build{
		AppName: "Echo-CA",
		Version: "v1beta",
	}

	server := fiber.New(fiber.Config{
		AppName: data.AppName,
	})
	server.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(data)
	})

	server.Use(logger.New())
	server.Listen(":8080")

}

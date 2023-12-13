package main

import (
	"alliance-calc-api/db"
	"alliance-calc-api/initializers"
	"alliance-calc-api/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	logs.Init()
	initializers.LoadEnvVariables()
	db.Initialize()
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋")
	})
}
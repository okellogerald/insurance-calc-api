package main

import (
	"insurance-calc-api/controllers/plans"
	"insurance-calc-api/db"
	"insurance-calc-api/env"
	"insurance-calc-api/logs"
	authMiddleware "insurance-calc-api/middleware/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	logs.Init()
	env.Load()
	db.Init()
	plans.Init()
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋")
	})

	api := app.Group("/api")
	api.Get("/", authMiddleware.ValidateSession, plans.GetPlans)

	err := app.Listen(":3300")
	if err != nil {
		logs.Error("server failed to start")
	}
}

// compiledaemon --command="./insurance-calc-api"

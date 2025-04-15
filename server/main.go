package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sagakyuuu/mahasiswa/server/config"
	"github.com/sagakyuuu/mahasiswa/server/router"
)

func main() {
	app := fiber.New()

	config.ConnectDB()

	app.Use(cors.New())

	api := app.Group("/api/v1")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Welcome to mahasiswa public API"})
	})
	router.AuthRouter(api)

	app.Listen(":5000")
}

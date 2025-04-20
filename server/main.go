package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()
	api := app.Group("/api/v1")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fmt.Println("Server is running on port", os.Getenv("PORT"))
	app.Listen(":" + os.Getenv("PORT"))

}

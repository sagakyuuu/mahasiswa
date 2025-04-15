package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sagakyuuu/mahasiswa/server/handlers"
)

func AuthRouter(app fiber.Router) {
	api := app.Group("/auth")
	api.Post("/register", handlers.Register)
}

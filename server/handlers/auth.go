package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sagakyuuu/mahasiswa/server/models"
)

var jwtSec = []byte("akjdja28jsbadkjab")

func Register(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	fmt.Println("Email:", user.Email)
	fmt.Println("Password:", user.Password)

	if user.Email == "" || user.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Email dan Password tidak boleh kosong"})
	}

	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(jwtSec)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal memuat token"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": signed,
		"email": user.Email,
	})
}


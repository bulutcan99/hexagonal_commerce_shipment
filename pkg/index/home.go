package index

import (
	"github.com/gofiber/fiber/v3"
)

func Homepage(c fiber.Ctx) error {
	return c.Render("index",
		fiber.Map{
			"Awesome": "Go-Commerce",
			"Home":    "active",
		})
}

func Signup(c fiber.Ctx) error {
	return c.Render("sign-up-screen",
		fiber.Map{
			"Awesome":  "Go-Commerce",
			"Register": "active",
		})
}

func Login(c fiber.Ctx) error {
	return c.Render("login-screen",
		fiber.Map{
			"Awesome": "Go-Commerce",
			"Login":   "active",
		})
}

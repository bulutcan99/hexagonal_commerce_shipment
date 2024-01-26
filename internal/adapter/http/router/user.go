package router

import (
	"github.com/gofiber/fiber/v3"
)

func UserRoute(r fiber.Router, user *controller.UserController) {
	route := r.Group("/v1/users")
	route.Post("/register", user.Register)
}

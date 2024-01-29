package router

import (
	"github.com/bulutcan99/commerce_shipment/internal/adapter/http/controller"
	"github.com/gofiber/fiber/v3"
)

func PermissionRoute(r fiber.Router, user *controller.UserController) {
	route := r.Group("/v1/users")
	route.Post("/register", user.Register)
}

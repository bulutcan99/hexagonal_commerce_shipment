package router

import (
	"github.com/bulutcan99/commerce_shipment/internal/adapter/http/controller"
	"github.com/gofiber/fiber/v3"
)

func AuthRoute(r fiber.Router, auth *controller.AuthController) {
	route := r.Group("/v1")
	route.Post("/register", auth.Register)
	route.Post("/login", auth.Login)
}

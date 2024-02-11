package router

import (
	fiber_go "github.com/bulutcan99/commerce_shipment/internal/adapter/fiber"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/http/controller"
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/gofiber/fiber/v3"
)

func UserRoute(r fiber.Router, token port.ITokenService, user *controller.UserController) {
	route := r.Group("/v1/user").Use(fiber_go.AuthMiddleware(token))
	route.Get("/get-users", user.GetAllUsers)
}

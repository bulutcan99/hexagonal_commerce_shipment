package router

import (
	"github.com/bulutcan99/commerce_shipment/internal/adapter/http/controller"
	"github.com/gofiber/fiber/v3"
)

func PermissionRoute(r fiber.Router, permission *controller.PermissionController) {
	// route := r.Group("/v1/permissions")
	// route.Post("/update/:user_id", permission.UpdatePermission)
}

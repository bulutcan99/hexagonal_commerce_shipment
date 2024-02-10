package controller

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
)

type AuthController struct {
	userService port.IAuthService
}

func NewUserController(userService port.IUserService) *UserController {
	return &UserController{
		userService,
	}
}

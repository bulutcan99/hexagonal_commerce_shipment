package controller

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
)

type UserController struct {
	userService port.IUserService
}

func NewUserController(userService port.IUserService) *UserController {
	return &UserController{
		userService,
	}
}

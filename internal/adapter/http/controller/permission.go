package controller

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
)

type PermissionController struct {
	permissionService port.IPermissionService
	userService       port.IUserService
}

func NewPermissionController(permissionService port.IPermissionService, userService port.IUserService) *PermissionController {
	return &PermissionController{
		permissionService,
		userService,
	}
}

type permissionReqBody struct {
	Entry     int  `json:"entry" binding:"required"`
	AddFlag   bool `json:"add_flag" binding:"required"`
	AdminFlag bool `json:"admin_flag" binding:"required"`
}

// func (u *PermissionController) UpdatePermission(c fiber.Ctx) error

package controller

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	userService port.IUserService
}

func NewUserController(userService port.IUserService) *UserController {
	return &UserController{
		userService,
	}
}

func (c *UserController) GetAllUsers(ctx fiber.Ctx) error {
	users, err := c.userService.ListAllUsers(ctx.Context())
	if err != nil {
		return ctx.Status(int(err.Code)).JSON(fiber.Map{"error": err.Message})
	}

	ctx.JSON(fiber.Map{"users": users})
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"users": users})
}

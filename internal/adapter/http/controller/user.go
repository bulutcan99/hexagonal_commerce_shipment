package controller

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/goccy/go-json"
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

type getUsersLimitReqBody struct {
	Skip  uint64 `json:"skip" binding:"required"`
	Limit uint64 `json:"limit" binding:"required"`
}

func (c *UserController) GetUsersLimit(ctx fiber.Ctx) error {
	var reqBody getUsersLimitReqBody
	body := ctx.Body()
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "error while trying to parse body",
		})
	}

	users, domainErr := c.userService.ListUsersWithLimit(ctx.Context(), reqBody.Skip, reqBody.Limit)
	if domainErr != nil {
		return ctx.Status(int(domainErr.Code)).JSON(
			fiber.Map{"error": domainErr.Message})
	}

	ctx.JSON(fiber.Map{"users": users})
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"users": users})
}

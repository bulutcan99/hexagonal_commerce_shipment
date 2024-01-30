package controller

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"log/slog"
)

type UserController struct {
	userService port.IUserService
}

func NewUserController(userService port.IUserService) *UserController {
	return &UserController{
		userService,
	}
}

type registerUserReqBody struct {
	Name               string `json:"name" binding:"required"`
	Surname            string `json:"surname" binding:"required"`
	Email              string `json:"email" binding:"required,email"`
	Password           string `json:"password" binding:"required,min=8"`
	Address            string `json:"address" binding:"required"`
	NotificationRadius uint64 `json:"notification_radius"`
}

func (u *UserController) Register(c fiber.Ctx) error {
	var reqBody registerUserReqBody
	body := c.Body()
	if err := json.Unmarshal(body, &reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "error while trying to parse body",
		})
	}

	user := domain.User{
		Name:               reqBody.Name,
		Surname:            reqBody.Surname,
		Email:              reqBody.Email,
		Password:           reqBody.Password,
		Address:            reqBody.Address,
		NotificationRadius: reqBody.NotificationRadius,
	}

	userData, err := u.userService.Register(c.Context(), &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "error while trying to register user: " + err.Message,
		})
	}
	slog.Info("User Registered Successfully! User:", userData)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "user registered successfully",
		"data":  userData,
	})
}

package controller

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"log/slog"
	"time"
)

type AuthController struct {
	authService port.IAuthService
}

func NewAuthController(authService port.IAuthService) *AuthController {
	return &AuthController{
		authService,
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

func (a *AuthController) Register(c fiber.Ctx) error {
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

	userData, err := a.authService.Register(c.Context(), &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "error while trying to register user: " + err.Message,
		})
	}

	slog.Info("User Registered Successfully! User:", userData.Email)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "user registered successfully",
		"data":  userData,
	})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *AuthController) Login(c fiber.Ctx) error {
	var reqBody LoginRequest
	body := c.Body()
	if err := json.Unmarshal(body, &reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "error while trying to parse body",
		})
	}

	token, err := a.authService.Login(c.Context(), reqBody.Email, reqBody.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "error while trying to login: " + err.Message,
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour),
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "login successful",
		"data":  token,
	})
}

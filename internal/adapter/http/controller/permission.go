package controller

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/domain"
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"log/slog"
	"strconv"
	"time"
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

func (u *PermissionController) AddPermission(c fiber.Ctx) error {
	userId := c.Params("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "user id must be a number",
		})
	}

	userData := u.userService.
	var reqBody permissionReqBody
	body := c.Body()
	if err := json.Unmarshal(body, &reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "error while trying to parse body",
		})
	}

	permissionData := domain.Permission{
		Entry:     reqBody.Entry,
		AddFlag:   reqBody.AddFlag,
		AdminFlag: reqBody.AdminFlag,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	permission, permissionErr := u.permissionService.AssignPermission(c.Context(), &permissionData, uint64(userIdInt))
	if permissionErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "error while trying to assign permission",
			"data":  err.Error(),
		})
	}

	slog.Info("Permission Assigned Successfully! Permission:", permission)
	return nil
}

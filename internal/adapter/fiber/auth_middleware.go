package fiber_go

import (
	"github.com/bulutcan99/commerce_shipment/internal/core/port"
	"github.com/gofiber/fiber/v3"
	"strings"
)

const (
	// AuthHeader is the name of the header that contains the access token
	AuthHeader = "authorization"
	// AuthType is the type of the authorization header
	AuthType = "bearer"
	// AuthPayloadKey is the key for the authorization payload in the context
	AuthPayload = "auth_payload"
)

// AuthMiddleware is a middleware to check if the user is authenticated
func AuthMiddleware(tokenService port.ITokenService) func(fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		authorizationHeader := ctx.Get(AuthHeader)

		if len(authorizationHeader) == 0 {
			return fiber.NewError(fiber.StatusUnauthorized, "empty authorization header")
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) != 2 {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid authorization header")
		}

		currentAuthorizationType := strings.ToLower(fields[0])
		if currentAuthorizationType != AuthType {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid authorization type")
		}

		accessToken := fields[1]
		payload, err := tokenService.VerifyToken(accessToken)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Message)
		}

		ctx.Locals(AuthPayload, payload)
		return ctx.Next()
	}
}

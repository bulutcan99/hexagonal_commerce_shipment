package route

import (
	"github.com/bulutcan99/commerce_shipment/pkg/index"
	"github.com/gofiber/fiber/v3"
)

func Index(b string, app *fiber.App) {
	app.Get(b, index.Homepage)
	app.Get(b+"auth/register", index.Signup)
	app.Get(b+"auth/login", index.Login)
}

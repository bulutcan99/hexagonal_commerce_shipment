package fiber_go

import (
	"github.com/bulutcan99/commerce_shipment/internal/adapter/env"
	go_json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"time"
)

var READ_TIMEOUT_SECONDS_COUNT = &env.Env.ServerReadTimeout

func ConfigFiber() fiber.Config {
	return fiber.Config{
		ReadTimeout:   time.Minute * time.Duration(*READ_TIMEOUT_SECONDS_COUNT),
		StrictRouting: false,
		CaseSensitive: false,
		BodyLimit:     16 * 1024 * 1024,
		JSONEncoder:   go_json.Marshal,
		JSONDecoder:   go_json.Unmarshal,
		AppName:       "Go-Commerce",
		Immutable:     true,
	}
}

package fiber_go

import (
	"context"
	"fmt"
	"github.com/bulutcan99/commerce_shipment/pkg/env"
	"github.com/gofiber/fiber/v3"
	"log"
)

var (
	Host = &env.Env.Host
	Port = &env.Env.ServerPort
)

func FiberListen(ctx context.Context, a *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		<-ctx.Done()
		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	fiberConnURL := fmt.Sprintf("%s:%d", *Host, *Port)
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

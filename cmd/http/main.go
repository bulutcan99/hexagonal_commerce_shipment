package http

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/cache/redis"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/fiber"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/fiber/route"
	repository "github.com/bulutcan99/commerce_shipment/internal/adapter/repository/postgres"
	"github.com/bulutcan99/commerce_shipment/pkg/env"
	"github.com/gofiber/fiber/v3"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var (
	slogger *slog.JSONHandler
	Env     *env.ENV
)

func Init() {
	Env = env.ParseEnv()
	slogger = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	logger := slog.New(slogger)
	slog.SetDefault(logger)
}

func main() {
	Init()
	slog.Info("Starting server...")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	db, err := repository.NewDB(ctx)
	if err != nil {
		slog.Error("Error connecting to database")
		panic(err)
	}
	defer db.Close()

	slog.Info("Database connected")
	cache, err := redis.NewRedisCache(ctx)
	if err != nil {
		slog.Error("Error connecting to redis")
		panic(err)
	}
	defer cache.Close()

	slog.Info("Redis connected")
	cfgFiber := fiber_go.ConfigFiber()
	app := fiber.New(cfgFiber)
	slog.Info("Fiber initialized")
	fiber_go.MiddlewareFiber(app)
	route.Index("/", app)
	fiber_go.FiberListen(ctx, app)
}

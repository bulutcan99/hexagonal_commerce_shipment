package http

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/config"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/env"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/fiber"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/fiber/route"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/storage/postgres"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/storage/redis"
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
	cfg := config.New()
	slog.Info("Config initialized")
	db, err := psql.NewDB(ctx, cfg.PSQL)
	if err != nil {
		slog.Error("Error connecting to database")
		panic(err)
	}
	defer db.Close()

	slog.Info("Database connected:", config.DbPort)
	err = db.Migrate()
	if err != nil {
		slog.Error("Error migrating database")
		panic(err)
	}

	slog.Info("Database migrated")
	cache, err := redis.NewRedisCache(ctx, cfg.Redis)
	if err != nil {
		slog.Error("Error connecting to redis")
		panic(err)
	}
	defer cache.Close()

	slog.Info("Redis connected:", config.RedisPort)
	cfgFiber := fiber_go.ConfigFiber()
	app := fiber.New(cfgFiber)
	slog.Info("Fiber initialized")
	fiber_go.MiddlewareFiber(app)
	route.Index("/", app)
	fiber_go.FiberListen(ctx, app)
}

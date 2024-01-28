package cmd

import (
	"context"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/config"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/env"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/fiber"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/http/controller"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/http/router"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/logger"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/storage/postgres"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/storage/postgres/repository"
	"github.com/bulutcan99/commerce_shipment/internal/adapter/storage/redis"
	"github.com/bulutcan99/commerce_shipment/internal/core/service"
	"github.com/gofiber/fiber/v3"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var (
	Env *env.ENV
)

func Init() {
	Env = env.ParseEnv()
	logger.Set()
}

func Run() {
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

	slog.Info("Database connected!")
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

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, cache)
	userHandler := controller.NewUserController(userService)
	slog.Info("Redis connected!")
	cfgFiber := fiber_go.ConfigFiber()
	app := fiber.New(cfgFiber)
	slog.Info("Fiber initialized")
	fiber_go.MiddlewareFiber(app)
	slog.Info("Fiber middleware initialized")
	router.UserRoute(app, userHandler)
	fiber_go.FiberListen(ctx, app)
}

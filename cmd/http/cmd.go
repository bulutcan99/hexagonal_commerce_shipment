package cmd

import (
	"context"
	paseto_token "github.com/bulutcan99/commerce_shipment/internal/adapter/auth/paseto"
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

	tokenService, err := paseto_token.New(cfg.Token)
	if err != nil {
		slog.Error("Error initializing token service")
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)

	authService := service.NewAuthService(userRepo, permissionRepo, cache, tokenService)
	userService := service.NewUserService(userRepo, permissionRepo, cache)
	permissionService := service.NewPermissionService(permissionRepo, cache)

	authController := controller.NewAuthController(authService)
	// userHandler := controller.NewUserController(userService)
	permissionHandler := controller.NewPermissionController(permissionService, userService)

	slog.Info("Redis connected!")
	cfgFiber := fiber_go.ConfigFiber()
	app := fiber.New(cfgFiber)
	slog.Info("Fiber initialized")
	fiber_go.MiddlewareFiber(app)
	slog.Info("Fiber middleware initialized")
	router.AuthRoute(app, authController)
	router.PermissionRoute(app, permissionHandler)
	fiber_go.FiberListen(ctx, app)
}

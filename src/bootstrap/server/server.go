package server

import (
	"context"
	"fmt"

	_ "github.com/erdemkosk/fiber-fx-boilerplate/docs"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/config"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/handler"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/middleware"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/repository"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/routes"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewFiberApp(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          middleware.ErrorHandler,
	})

	app.Use(middleware.RateLimiterMiddleware(cfg))
	app.Use(middleware.CorsMiddleware())
	if cfg.App.SwaggerEnabled {
		app.Get("/swagger/*", middleware.SwaggerMiddleware())
	}

	return app
}

func StartFiberApp(lifecycle fx.Lifecycle, app *fiber.App, logger *zap.Logger, cfg *config.Config) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info(fmt.Sprintf("🚀 Fiber server starting on :%d", cfg.App.Port))

				if cfg.App.SwaggerEnabled {
					logger.Info("📚 Swagger documentation is available at: /swagger")
				}

				if err := app.Listen(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
					logger.Fatal("Fiber server failed to start", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("🛑 Shutting down Fiber server")
			return app.Shutdown()
		},
	})
}

var Module = fx.Options(
	repository.Module,
	handler.Module,
	service.Module,
	routes.Module,
	fx.Provide(NewFiberApp),
	fx.Invoke(StartFiberApp),
)

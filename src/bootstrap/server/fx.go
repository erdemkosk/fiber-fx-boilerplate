package server

import (
	"context"
	"fmt"

	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/config"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/handler"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/repository"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/routes"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		DisableStartupMessage: true,
		EnablePrintRoutes:     false,
	})
}

func StartFiberApp(lifecycle fx.Lifecycle, app *fiber.App, logger *zap.Logger, cfg *config.Config) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info(fmt.Sprintf("🚀 Fiber server starting on :%d", cfg.App.Port))
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

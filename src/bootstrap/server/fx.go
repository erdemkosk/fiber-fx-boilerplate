package server

import (
	"context"

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

func StartFiberApp(lifecycle fx.Lifecycle, app *fiber.App, logger *zap.Logger) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("🚀 Fiber server starting on :3000")
				if err := app.Listen(":3000"); err != nil {
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

package main

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/config"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/logger"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/mongodb"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		config.Module,
		logger.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.NopLogger
		}),
		mongodb.Module,
		server.Module,
	)
	app.Run()
}

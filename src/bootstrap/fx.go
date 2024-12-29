package bootstrap

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/config"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/logger"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/mongodb"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Options(
	config.Module,
	logger.Module,
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log.Named("fx").WithOptions(zap.IncreaseLevel(zap.ErrorLevel))}
	}),
	mongodb.Module,
	server.Module,
)

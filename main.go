package main

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/bootstrap/config"
	"github.com/erdemkosk/fiber-fx-boilerplate/bootstrap/mongodb"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		mongodb.Module,
		fx.Invoke(startApp),
	)
	app.Run()
}

func startApp(cfg *config.Config, db *mongodb.MongoDB) {

}

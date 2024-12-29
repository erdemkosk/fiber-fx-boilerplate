package main

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}

package main

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap"
	"go.uber.org/fx"
)

// @title Fiber FX Boilerplate API
// @version 1.0
// @description This is a sample server for Fiber Boilerplate.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email erdemkosk@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8090
// @BasePath /api
// @schemes http https
func main() {
	fx.New(bootstrap.Module).Run()
}

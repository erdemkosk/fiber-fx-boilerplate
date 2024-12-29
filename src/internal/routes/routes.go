package routes

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/handler"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Route struct {
	Fiber *fiber.App
}

func NewRoute(app *fiber.App) Route {
	return Route{
		Fiber: app,
	}
}

var Module = fx.Options(
	fx.Provide(NewRoute),
	fx.Invoke(registerRoutes),
)

func registerRoutes(
	fooHandler *handler.FooHandler,
	route Route,
) {
	api := route.Fiber.Group("/api")

	fooRouter := NewFooRouter(fooHandler)
	fooRouter.SetupRoutes(api)
}

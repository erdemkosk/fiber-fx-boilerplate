package routes

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/handler"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/middleware"
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/model"
	"github.com/gofiber/fiber/v2"
)

type FooRouter struct {
	handler *handler.FooHandler
}

func NewFooRouter(handler *handler.FooHandler) *FooRouter {
	return &FooRouter{
		handler: handler,
	}
}

func (r *FooRouter) SetupRoutes(router fiber.Router) {
	foo := router.Group("/foo")
	foo.Get("/", r.handler.GetAll)
	foo.Get("/:id", r.handler.GetById)
	foo.Post("/", middleware.RequestValidator(model.Foo{}), r.handler.Create)
	foo.Put("/:id", middleware.RequestValidator(model.Foo{}), r.handler.Update)
}

package routes

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/handler"
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
	foo.Post("/", r.handler.Create)
	foo.Put("/:id", r.handler.Update)
}

package routes

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type HealthRouter struct {
	handler *handler.HealthHandler
}

func NewHealthRouter(handler *handler.HealthHandler) *HealthRouter {
	return &HealthRouter{
		handler: handler,
	}
}

func (r *HealthRouter) SetupRoutes(router fiber.Router) {
	router.Get("/health", r.handler.Check)
}

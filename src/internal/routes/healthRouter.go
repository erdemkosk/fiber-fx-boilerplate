package routes

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type HealthRouter struct {
	handler handler.IHealthHandler
}

func NewHealthRouter(handler handler.IHealthHandler) *HealthRouter {
	return &HealthRouter{
		handler: handler,
	}
}

func (r *HealthRouter) SetupRoutes(router fiber.Router) {
	router.Get("/health", r.handler.Check)
}

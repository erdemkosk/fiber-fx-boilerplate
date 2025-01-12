package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type IHealthHandler interface {
	Check(c *fiber.Ctx) error
}

type HealthHandler struct{}

var _ IHealthHandler = (*HealthHandler)(nil)

func NewHealthHandler() IHealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"time":   time.Now(),
		"status": "ok",
	})
}

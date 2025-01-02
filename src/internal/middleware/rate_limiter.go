package middleware

import (
	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiterMiddleware(cfg *config.Config) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        cfg.App.RateLimit,
		Expiration: cfg.App.RateWindow,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded",
			})
		},
	})
}

package middleware

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/response"
)

func RateLimit(max, expirationSecond int, message *string) fiber.Handler {
	if message == nil {
		defaultMessage := "too many request, please wait."
		message = &defaultMessage
	}
	return limiter.New(
		limiter.Config{
			Max:               max,
			Expiration:        time.Duration(expirationSecond) * time.Second,
			LimiterMiddleware: limiter.SlidingWindow{}, // sliding window rate limiter,
			LimitReached: func(c *fiber.Ctx) error {
				return response.ErrorMessage(c, fiber.StatusTooManyRequests, *message, errors.New(*message))

			},
		},
	)
}

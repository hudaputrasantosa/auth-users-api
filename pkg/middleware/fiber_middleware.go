package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(app *fiber.App){
	corsConfig := cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}

	limiterConfig := limiter.Config{
		Max:               10,
		Expiration:        1 * time.Minute,
		LimiterMiddleware: limiter.SlidingWindow{}, // sliding window rate limiter,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"errors":  true,
				"message": "Too many requests, please try again later.",
			})
		},
	}

	swaggerConfig := swagger.Config{
	BasePath: "/",
    FilePath: "./docs/swagger.yaml",
    Path:     "swagger",
    Title:    "Swagger API Docs",
	}

	app.Use(
	helmet.New(),
	logger.New(),
	cors.New(corsConfig),
	limiter.New(limiterConfig),
	recover.New(), // recover will catch panics like from handler and recover the panic and throw to fiber error handler
	swagger.New(swaggerConfig),
	)
}

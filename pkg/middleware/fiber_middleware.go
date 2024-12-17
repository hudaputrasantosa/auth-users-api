package middleware

import (
	"time"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
)

// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(app *fiber.App) {
	corsConfig := cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}

	// disabled csrf middleware
	// csrfConfig := csrf.Config{
	// 	KeyLookup:      "header:X-Csrf-Token", // string in the form of '<source>:<key>' that is used to extract token from the request
	// 	CookieName:     "csrf_",            // name of the session cookie
	// 	CookieSameSite: "Strict",              // indicates if CSRF cookie is requested by SameSite
	// 	Expiration:     3 * time.Hour,         // expiration is the duration before CSRF token will expire
	// }

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
		// csrf.New(csrfConfig),
		limiter.New(limiterConfig),
		recover.New(), // recover will catch panics like from handler and recover the panic and throw to fiber error handler
		swagger.New(swaggerConfig),
	)
}

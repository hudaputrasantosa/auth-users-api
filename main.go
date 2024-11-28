package main

import (
	// "fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/cors"
	appLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database"
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database/migration"
	"github.com/hudaputrasantosa/auth-users-api/internal/interface/http/route"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"

	_ "github.com/lib/pq"
)

func main() {
	//initial database
	database.Connect()
	//migration database
	migration.Migration()
	// initial fiber
	app := fiber.New()

	app.Use(helmet.New())
	app.Use(appLogger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))
	app.Use(limiter.New(limiter.Config{
		Max:               10,
		Expiration:        1 * time.Minute,
		LimiterMiddleware: limiter.SlidingWindow{}, // sliding window rate limiter,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"errors":  true,
				"message": "Too many requests, please try again later.",
			})
		},
	}))
	app.Use(recover.New()) // recover will catch panics like from handler and recover the panic and throw to fiber error handler

	app.Use(swagger.New(swagger.Config{
	BasePath: "/",
    FilePath: "./docs/swagger.yaml",
    Path:     "swagger",
    Title:    "Swagger API Docs",
	}))

	route.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	logger.Info("Server Running ..")
	app.Listen(":8080")
}

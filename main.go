package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database"
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database/migration"
	"github.com/hudaputrasantosa/auth-users-api/internal/interface/http/routes"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
	"github.com/hudaputrasantosa/auth-users-api/pkg/middleware"

	_ "github.com/lib/pq"
)

func main() {
	//initial database
	database.Connect()
	//migration database
	migration.Migration()
	// initial fiber
	app := fiber.New()
	// Middleware application
	middleware.FiberMiddleware(app)
	// Main router
	routes.SetupRoutes(app)

	logger.Info("Server Running ..")
	// Start server on port 8080
	app.Listen(":8080")
}

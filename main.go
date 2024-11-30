package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database"
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database/migration"
	"github.com/hudaputrasantosa/auth-users-api/internal/interface/http/routes"
	"github.com/hudaputrasantosa/auth-users-api/pkg/middleware"
	"github.com/hudaputrasantosa/auth-users-api/internal/config"
	"github.com/hudaputrasantosa/auth-users-api/pkg/server"
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
	// Start server
	appEnv := config.Config("APP_ENV")
	if appEnv == "development" || appEnv == "staging" {
		server.StartServer(app)
	} else {
		server.StartServerWithGracefulShutdown(app)
	}
}

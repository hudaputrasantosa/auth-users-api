package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"github.com/hudaputrasantosa/auth-users-api/internal/config"
	"github.com/hudaputrasantosa/auth-users-api/internal/delivery"
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database"
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database/migration"
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/server"
	"github.com/hudaputrasantosa/auth-users-api/pkg/middleware"
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
	//Repositories
	repositories := delivery.NewRepository(database.DB.Db)
	// Services
	services := delivery.NewService(delivery.Inject{
		Repository: repositories,
	})
	// Rest API Handler
	restHandler := delivery.NewRestHandler(services.UserService, services.AuthService)

	app = restHandler.Init()

	// Start server
	appEnv := config.Config("APP_ENV")
	if appEnv == "development" || appEnv == "staging" {
		server.StartServer(app)
	} else {
		server.StartServerWithGracefulShutdown(app)
	}
}

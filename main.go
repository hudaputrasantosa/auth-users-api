package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"github.com/hudaputrasantosa/auth-users-api/internal/config"
	"github.com/hudaputrasantosa/auth-users-api/internal/delivery"
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database"
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database/migration"
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/server"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
	"github.com/hudaputrasantosa/auth-users-api/pkg/middleware"
)

func main() {
	defer logger.Log.Sync() // Pastikan log disinkronisasi sebelum aplikasi berakhir
	defer config.InitProvider()

	//initial database
	database.Connect()
	//migration database
	migration.Migration()
	//Initiate Redis
	config.InitRedis()
	// initial fiber
	app := fiber.New()
	// Middleware application
	middleware.FiberMiddleware(app)

	// Setup Cloudinary client
	// cld, err := config.SetupCloudinary()
	// if err != nil {
	// 	log.Fatalf("Failed to setup Cloudinary: %v", err)
	// }
	//Repositories
	repositories := delivery.NewRepository(database.DB.Db)
	// Services
	services := delivery.NewService(delivery.Inject{
		Repository: repositories,
		Redis:      config.RedisClient,
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

package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudaputrasantosa/auth-users-api/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	authRoute := api.Group("/auth")
	userRoute := api.Group("/users")

	// Handle Authentication
	authRoute.Post("/register", handler.CreateUser)
	authRoute.Get("/", handler.GetAllUser)

	// Handle Users
	userRoute.Get("/", handler.GetAllUser)
}

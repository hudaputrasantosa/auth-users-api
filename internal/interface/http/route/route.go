package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hudaputrasantosa/auth-users-api/internal/interface/http/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	// Grouping routes
	authRoute := api.Group("/auth")
	userRoute := api.Group("/users")

	// Handle Authentication
	authRoute.Post("/register", controllers.CreateUser)
	authRoute.Post("/login", controllers.ValidateUser)

	// Handle Users Management
	userRoute.Post("/", controllers.CreateUser)
	userRoute.Get("/", controllers.FindUsers)
	userRoute.Get("/pagination", controllers.FindUsersPagination)
	userRoute.Get("/:id", controllers.FindUserById)
	userRoute.Delete("/:id", controllers.DeleteUserById)
}

package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hudaputrasantosa/auth-users-api/internal/interface/http/routes/auth"
	"github.com/hudaputrasantosa/auth-users-api/internal/interface/http/routes/user"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/response"
)


func SetupRoutes(app *fiber.App) {
	app.Get("health-checker", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Golang, Fiber, and GORM")
	})

	auth.AuthRoutes(app)
	user.UserRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return response.ErrorMessage(c, fiber.StatusNotFound, "Route Not found", nil)
	})
}

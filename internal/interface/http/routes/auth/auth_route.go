package auth

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hudaputrasantosa/auth-users-api/internal/interface/http/controllers"
)

func AuthRoutes(app *fiber.App){
	api := app.Group("/api")
	v1 := api.Group("/v1/auth")

	//Routes Version 1.0
	v1.Post("/register", controllers.CreateUser)
	v1.Post("/login", controllers.ValidateUser)

}

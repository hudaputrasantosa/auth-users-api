package user

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hudaputrasantosa/auth-users-api/internal/interface/http/controllers"
)

func UserRoutes(app *fiber.App){
	api := app.Group("/api")
	v1 := api.Group("/v1/users")

	//Routes Version 1.0
	v1.Post("/", controllers.CreateUser)
	v1.Get("/", controllers.FindUsers)
	v1.Get("/pagination", controllers.FindUsersPagination)
	v1.Get("/:id", controllers.FindUserById)
	v1.Delete("/:id", controllers.DeleteUserById)

}

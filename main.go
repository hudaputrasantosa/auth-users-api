package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hudaputrasantosa/auth-users-api/database"
	"github.com/hudaputrasantosa/auth-users-api/database/migration"
	"github.com/hudaputrasantosa/auth-users-api/router"
	_ "github.com/lib/pq"
)

func main() {
	//initial database
	database.Connect()
	//migration database
	migration.Migration()
	// initial fiber
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	fmt.Println("Server Running ..")
	app.Listen(":8080")
}

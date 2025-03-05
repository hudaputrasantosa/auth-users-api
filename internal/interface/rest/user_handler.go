package rest

import (
	"github.com/gofiber/fiber/v2"

	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	service "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/services"
	"github.com/hudaputrasantosa/auth-users-api/internal/middleware"
	middle "github.com/hudaputrasantosa/auth-users-api/pkg/middleware"
)

type handleUser struct {
	userClient service.UserService
}

// Router Users
func NewHandleUserRoute(userClient service.UserService, router *fiber.App) {
	handlerUser := handleUser{
		userClient,
	}

	api := router.Group("/api")

	// Admin Router Group and Middleware  Version 1.0
	adminRouterV1 := api.Group("/v1/admin")
	adminRouterV1.Use(middleware.Authorized(), middle.RateLimit(30, 60, nil))

	// User Management Route
	adminRouterV1.Use(middleware.RoleProtection(model.Admin, handlerUser.userClient))
	adminRouterV1.Post("/", handlerUser.createUser)
	adminRouterV1.Get("/", handlerUser.findUsers)
	adminRouterV1.Get("/pagination", handlerUser.findUsersPagination)
	adminRouterV1.Get("/:id", handlerUser.findUserById)
	adminRouterV1.Delete("/:id", handlerUser.deleteUserById)

	// ======================================================================================

	// Member Router Group and Middleware  Version 1.0
	memberRouterV1 := api.Group("/v1/member")
	memberRouterV1.Use(middleware.Authorized(), middle.RateLimit(30, 60, nil))
}

package delivery

import (
	"github.com/gofiber/fiber/v2"
	authService "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/services"
	userService "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/services"
	"github.com/hudaputrasantosa/auth-users-api/internal/interface/rest"
)

type handler struct {
	user userService.UserService
	auth authService.AuthService
}

func NewRestHandler(
	user userService.UserService,
	auth authService.AuthService,
) *handler {
	return &handler{
		user,
		auth,
	}
}

func (h *handler) Init() *fiber.App {
	router := fiber.New()

	router.Get("/health-checker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("Hello World")
	})

	rest.NewHandleAuthRoute(h.auth, router)
	rest.NewHandleUserRoute(h.user, router)

	return router
}

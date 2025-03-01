package middleware

import (
	"github.com/gofiber/fiber/v2"

	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	service "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/services"

	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/response"
)

func RoleProtection(role model.Role, userService service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		userId := c.Locals("user")

		// find user and get role
		user, status, err := userService.FindByID(ctx, userId.(string))
		if err != nil || user == nil {
			return response.ErrorMessage(c, status, "Failed to find user", err)
		}

		//check and throw if role invalid
		if user.Role != role {
			return response.ErrorMessage(c, fiber.StatusForbidden, "access not permitted", nil)
		}

		return c.Next()
	}
}

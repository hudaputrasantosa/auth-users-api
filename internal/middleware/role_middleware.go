package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	service "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/services"

	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/response"
)

func IsAdmin(userService service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		userId := c.Locals("user")
		// find user and get role
		user, status, err := userService.FindByID(ctx, userId.(string))
		if err != nil || user == nil {
			return response.ErrorMessage(c, status, "Failed to find user", err)
		}
		//check and throw if role not admin
		if user.Role != model.Admin {
			return response.ErrorMessage(c, fiber.StatusForbidden, "Forbidden: role is not admin", nil)
		}

		return c.Next()
	}

}

func IsMember() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		fmt.Println("role middleware")
		fmt.Println(user)
		// find user and get role

		//check and throw if role not admin
		if user != model.Member {
			return response.ErrorMessage(c, fiber.StatusForbidden, "Forbidden: role is not member", nil)
		}

		return c.Next()
	}
}

package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/response"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/token"
)

func Authorized() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.ErrorMessage(c, fiber.StatusUnauthorized, "Unauthorized: No authorization header", nil)
		}

		// Pastikan header menggunakan skema "Bearer"
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return response.ErrorMessage(c, fiber.StatusUnauthorized, "Invalid Authorization header forma", nil)
		}

		// Ambil token setelah "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, token.JwtKeyFunc)
		// Jika parsing gagal, atau token tidak valid
		fmt.Println(err)
		if err != nil {
			return response.ErrorMessage(c, fiber.StatusUnauthorized, "Invalid or expired token", nil)
		}

		// Ambil klaim dari token (opsional: jika Anda menggunakan custom claims)
		if claims, ok := token.Claims.(jwt.MapClaims)["id"]; ok {
			// Simpan data user dari klaim ke dalam context (opsional)
			c.Locals("user", claims)
		} else {
			return response.ErrorMessage(c, fiber.StatusUnauthorized, "Invalid token claims", nil)
		}

		// versi lain, fungsi mengambil id, kemudian mencari user by id dnegan saervice dan membuat object baru
		// 	id := decode_token.Claims.(jwt.MapClaims)["id"].(float64)

		// user_service := service.NewUserService(repository.NewUserRepository(), database.DB)
		// user, err := user_service.FindById(c.Context(), int(id))
		// if err != nil {
		// 	return helper.RespondError(c, fiber.StatusUnauthorized, "Unauthorized")
		// }

		// userSession := dto.UserSession{
		// 	Id:       user.Id,
		// 	Username: user.Username,
		// 	Role:     user.Role,
		// }

		// Lanjutkan ke handler berikutnya
		return c.Next()
	}

}

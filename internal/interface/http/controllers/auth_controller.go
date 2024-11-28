package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database"
	"github.com/hudaputrasantosa/auth-users-api/internal/interface/http/dto"
	"github.com/hudaputrasantosa/auth-users-api/internal/domain/entity"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/hash"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/response"
)

func ValidateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	var payload *dto.ValidateUserSchema;
	var user *entity.User

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, "Failed parsing", err)
	}

	// check email exist
	userResult := db.First(&user, "email = ?", payload.Email)
	if err := userResult.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.ErrorMessage(c, fiber.StatusNotFound, "Email not found or not registered", err)
		}
			return response.ErrorMessage(c, fiber.StatusBadGateway, err.Error(), err)
	}

	//chack password
	isPassword := hash.CheckPasswordHash(payload.Password, user.Password)
	if !isPassword {
		return response.ErrorMessage(c, fiber.StatusBadRequest, "Password is wrong", nil)
	}

	// check status user
	if !user.IsActive {
		// send email service to verification

		return response.ErrorMessage(c, fiber.StatusBadRequest, "User not active, Please to verification. check your email", nil)
	}

	// generate token jwt

	return response.SuccessMessage(c, fiber.StatusOK, "Success login")
}

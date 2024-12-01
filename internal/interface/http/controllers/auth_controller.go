package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database"
	"github.com/hudaputrasantosa/auth-users-api/internal/interface/http/dto"
	"github.com/hudaputrasantosa/auth-users-api/internal/domain/entity"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/hash"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/response"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/token"
)

func ValidateUser(c *fiber.Ctx) error {
	// DB connection
	db := database.DB.Db

	// Create or initial login struct payload
	var payload *dto.ValidateUserSchema;

	// Create user model to store data
	var user *entity.User

	// Check received data from JSON body.
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

	// generate new access token and refresh token jwt
	userToken, err := token.GenerateNewToken(user.ID.String())
	if err!= nil {
        return response.ErrorMessage(c, fiber.StatusInternalServerError, "Failed to generate token", err)
    }

	return response.SuccessMessageWithData(c, fiber.StatusOK, "Success login", userToken)
}

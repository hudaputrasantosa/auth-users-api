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
    "github.com/hudaputrasantosa/auth-users-api/pkg/helper/validation"
)

type UserTokenResponse struct {
    Role        entity.Role
    AccessToken string
    RefreshToken string
}

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

    // Validate data before proceessing.
    	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
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

    res := UserTokenResponse{
        Role:        user.Role,
        AccessToken: userToken.AccessToken,
        RefreshToken: userToken.RefreshToken,
    }

	return response.SuccessMessageWithData(c, fiber.StatusOK, "Success login", res)
}

func RegisterUser(c *fiber.Ctx) error {
	// DB connection
    db := database.DB.Db

    // Create or initial user struct payload
    var payload *dto.RegisterUserSchema;

    // Create user model to store data
    var user *entity.User

    // Check received data from JSON body.
    if err := c.BodyParser(&payload); err!= nil {
        return response.ErrorMessage(c, fiber.StatusBadRequest, "Failed parsing", err)
    }

	// check email validation

    // check email existing
    userResult := db.First(&user, "email =?", payload.Email)
    if userResult.Error == nil {
        return response.ErrorMessage(c, fiber.StatusConflict, "Email already registered", nil)
    }

	// if admin role, get header secret key for create admin user

    // hash password
    hashedPassword, err := hash.HashPassword(payload.Password)
	if err!= nil {
        return response.ErrorMessage(c, fiber.StatusInternalServerError, "Failed to hash password", err)
    }

	// generate otp token from jwt

	// sent otp to active email that registered

	// return success with token otp
}

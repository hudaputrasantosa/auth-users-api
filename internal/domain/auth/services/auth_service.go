package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/dtos"
	"github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/utils"
	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	globalUtils "github.com/hudaputrasantosa/auth-users-api/internal/utils"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/hash"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/token"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserTokenResponse struct {
	Role         model.Role
	AccessToken  string
	RefreshToken string
}

func (s *serviceAuth) ValidateUser(ctx context.Context, payload dto.ValidateUserSchema) (*UserTokenResponse, int, error) {
	// Create user model to store data
	var user *model.User

	// check email exist
	user, err := s.userRepository.FindByEmail(ctx, payload.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error("email not found or not registered")
			return nil, fiber.StatusNotFound, utils.FailedUserLogin
		}

		logger.Error(err.Error())
		return nil, fiber.StatusInternalServerError, err
	}

	//chack password
	isPassword := hash.CheckPasswordHash(payload.Password, user.Password)
	if !isPassword {
		logger.Error("password is not match")
		return nil, fiber.StatusConflict, utils.FailedUserLogin
	}

	// check status user
	if !user.IsActive {
		// send email service to verification [PLAN]
		fmt.Println("Sent email to verification")

		logger.Error("User not active, Please to verification. check your email")
		return nil, fiber.StatusBadRequest, errors.New("please to verification email")
	}

	// generate new access token and refresh token jwt
	userToken, err := token.GenerateNewToken(user.ID.String(), string(user.Role))
	if err != nil {
		logger.Error("failed generate token", zap.Error(err))
		return nil, fiber.StatusInternalServerError, globalUtils.ErrorGlobalPublicMessage
	}

	res := &UserTokenResponse{
		Role:         user.Role,
		AccessToken:  userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
	}

	return res, fiber.StatusOK, err
}

func (s *serviceAuth) RegisterUser(ctx context.Context, payload dto.RegisterUserSchema) (interface{}, int, error) {
	// check email existing
	user, err := s.userRepository.FindByEmail(ctx, payload.Email)
	if err != nil {
		logger.Error("Error register", zap.Error(err))
		return nil, fiber.StatusInternalServerError, utils.ErrorUserRegister
	} else if user != nil {
		return nil, fiber.StatusConflict, utils.FailedUserRegister
	}

	// if admin role, get header secret key for create admin user

	res, err := s.authRepository.Register(ctx, &payload)
	if err != nil {
		logger.Error("Error register", zap.Error(err))
		return nil, fiber.StatusInternalServerError, utils.ErrorUserRegister
	}

	// generate otp token from jwt
	// otpToken, err := token.GenerateNewToken(user.Email)
	// sent otp to active email that registered

	// return success with token otp
	return res, fiber.StatusCreated, nil
}

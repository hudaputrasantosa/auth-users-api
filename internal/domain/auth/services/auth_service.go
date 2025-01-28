package services

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/dtos"
	"github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/utils"
	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	globalUtils "github.com/hudaputrasantosa/auth-users-api/internal/utils"
	"github.com/hudaputrasantosa/auth-users-api/pkg/hash"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
	"github.com/hudaputrasantosa/auth-users-api/pkg/notification"
	"github.com/hudaputrasantosa/auth-users-api/pkg/token"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/templates"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserTokenResponse struct {
	Role         model.Role
	AccessToken  string
	RefreshToken string
}

type UserRegisterResponse struct {
	Email   string
	Token   string
	Expired string
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
		// send email service to verification
		// otpToken, err := token.GenerateNewToken(user.Email)
		otpToken := "12345"
		_, err = notification.MailersendNotification(&notification.RecipientInformation{
			Email: user.Email,
			Name:  user.Name,
		}, &templates.DataBodyInformation{
			Name:            user.Name,
			Otp:             otpToken,
			MessageTemplate: templates.Otp_template,
		})
		if err != nil {
			logger.Error("Failed send notification")
		}

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
		if err != gorm.ErrRecordNotFound {
			logger.Error("Error register", zap.Error(err))
			return nil, fiber.StatusInternalServerError, utils.ErrorUserRegister
		}
	} else if user != nil {
		return nil, fiber.StatusConflict, utils.FailedUserRegister
	}

	// if admin role, get header secret key for create admin user
	now := time.Now()

	hashPassword, err := hash.HashPassword(payload.Password)
	if err != nil {
		logger.Error("Error, failed to hash password", zap.Error(err))
		return nil, fiber.StatusInternalServerError, utils.ErrorUserRegister
	}

	createUser := model.User{
		ID:        uuid.New(),
		Name:      payload.Name,
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  hashPassword,
		Phone:     payload.Phone,
		CreatedAt: now,
		UpdatedAt: now,
	}

	user, err = s.authRepository.Register(ctx, createUser)
	if err != nil {
		logger.Error("Error register", zap.Error(err))
		return nil, fiber.StatusInternalServerError, utils.ErrorUserRegister
	}

	// generate otp token from jwt
	// otpToken, err := token.GenerateNewToken(user.Email)
	otpToken := "12345"

	// sent otp to active email that registered
	_, err = notification.MailersendNotification(&notification.RecipientInformation{
		Email: user.Email,
		Name:  user.Name,
	}, &templates.DataBodyInformation{
		Name:            user.Name,
		Otp:             otpToken,
		MessageTemplate: templates.Otp_template,
	})
	if err != nil {
		logger.Error("Failed send notification")
	}

	res := &UserRegisterResponse{
		Email:   user.Email,
		Token:   "token-dummy",
		Expired: "5minutes",
	}

	// return success with token otp
	return res, fiber.StatusCreated, nil
}

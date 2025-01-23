package services

import (
	"context"
	"errors"
	"fmt"

	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/dtos"
	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/hash"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/token"
	"gorm.io/gorm"
)

type UserTokenResponse struct {
	Role         model.Role
	AccessToken  string
	RefreshToken string
}

func (s *serviceAuth) ValidateUser(ctx context.Context, payload dto.ValidateUserSchema) (*UserTokenResponse, error) {
	// Create user model to store data
	var user *model.User

	// check email exist // change to service
	user, err := s.userRepository.FindByEmail(ctx, payload.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// give a log err with message "Email not found or not registered"
			return nil, errors.New("email or password is wrong")
			// return response.ErrorMessage(c, fiber.StatusNotFound, "Email not found or not registered", err)
		}
		return nil, err
		// return response.ErrorMessage(c, fiber.StatusBadGateway, err.Error(), err)
	}

	//chack password
	isPassword := hash.CheckPasswordHash(payload.Password, user.Password)
	if !isPassword {
		// give a log err with message "Email or password is wrong"
		return nil, errors.New("email or password is wrong")
		// return response.ErrorMessage(c, fiber.StatusBadRequest, "Password is wrong", nil)
	}

	// check status user
	if !user.IsActive {
		// send email service to verification [PLAN]
		fmt.Println("Sent email to verification")

		// give a log err with message "User not active, Please to verification. check your email"
		return nil, errors.New("please to verification email")
		// return response.ErrorMessage(c, fiber.StatusBadRequest, "User not active, Please to verification. check your email", nil)
	}

	// generate new access token and refresh token jwt
	userToken, err := token.GenerateNewToken(user.ID.String(), string(user.Role))
	if err != nil {
		// global error
		return nil, errors.New("ops! Something went wrong")
		// return response.ErrorMessage(c, fiber.StatusInternalServerError, "Failed to generate token", err)
	}

	res := &UserTokenResponse{
		Role:         user.Role,
		AccessToken:  userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
	}

	return res, err
}

func (s *serviceAuth) RegisterUser(ctx context.Context, payload dto.RegisterUserSchema) (interface{}, error) {
	// check email existing
	userResult, err := s.userRepository.FindByEmail(ctx, payload.Email)
	if err != nil {
		err := errors.New("error, can't process register")
		return nil, err
		// return response.ErrorMessage(c, fiber.StatusConflict, "Email already registered", nil)
	} else if userResult != nil {
		// give a log error
		// Email already registered
		err := errors.New("error, can't process register")
		return nil, err
	}

	// if admin role, get header secret key for create admin user

	res, err := s.authRepository.Register(ctx, &payload)
	if err != nil {
		// give a log error
		// Failed to create user
		err := errors.New("error, can't process register")
		return nil, err
		// return response.ErrorMessage(c, fiber.StatusInternalServerError, "Failed to create user", err)
	}

	// generate otp token from jwt
	// otpToken, err := token.GenerateNewToken(user.Email)
	// sent otp to active email that registered

	// return success with token otp
	return res, nil
	// return response.SuccessMessageWithData(c, fiber.StatusOK, "Success Register", user)
}

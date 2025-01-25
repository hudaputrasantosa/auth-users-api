package services

import (
	"context"

	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/dtos"
	authRepository "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/repositories"
	userRepository "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/repositories"
)

type serviceAuth struct {
	authRepository authRepository.AuthRepository
	userRepository userRepository.UserRepository
}

func NewAuthService(
	authRepository authRepository.AuthRepository,
	userRepository userRepository.UserRepository,
) *serviceAuth {
	return &serviceAuth{
		authRepository,
		userRepository,
	}
}

// Interface Auth Service untuk mengetahui beberapa schema header yang tersedia pada User service
type AuthService interface {
	ValidateUser(ctx context.Context, payload dto.ValidateUserSchema) (*UserTokenResponse, int, error)
	RegisterUser(ctx context.Context, payload dto.RegisterUserSchema) (interface{}, int, error)
	// Verification(ctx context.Context, otp string) (*model.User, error)
	// ResendVerification(ctx context.Context, token string) (*model.User, error)
	// ForgotPassword(ctx context.Context, contact *model.User) (*model.User, error)
	// ResendForgotPassword(ctx context.Context, token *model.User) (*model.User, error)
}

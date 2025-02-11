package services

import (
	"context"

	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/dtos"
	authRepository "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/repositories"
	userRepository "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/repositories"
	"github.com/redis/go-redis/v9"
)

type serviceAuth struct {
	authRepository authRepository.AuthRepository
	userRepository userRepository.UserRepository
	redis          *redis.Client
}

func NewAuthService(
	authRepository authRepository.AuthRepository,
	userRepository userRepository.UserRepository,
	redis *redis.Client,
) *serviceAuth {
	return &serviceAuth{
		authRepository,
		userRepository,
		redis,
	}
}

// Interface Auth Service untuk mengetahui beberapa schema header yang tersedia pada User service
type AuthService interface {
	ValidateUser(ctx context.Context, payload dto.ValidateUserSchema) (*UserTokenResponse, int, error)
	RegisterUser(ctx context.Context, payload dto.RegisterUserSchema) (interface{}, int, error)
	VerificationUser(ctx context.Context, payload dto.VerificationUser) (string, int, error)
	ResendVerificationUser(ctx context.Context, payload dto.ResendVerificationUser) (interface{}, int, error)
	// ForgotPassword(ctx context.Context, contact *model.User) (*model.User, error)
	// ResendForgotPassword(ctx context.Context, token *model.User) (*model.User, error)
}

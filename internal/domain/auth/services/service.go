package services

import (
	"context"

	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/dtos"
	authRepository "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/repositories"
	"github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
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
	// Member Section
	ValidateUser(ctx context.Context, payload dto.ValidateUserSchema, activitySource *models.UsersActivityHistory) (*UserTokenResponse, int, error)
	RegisterUser(ctx context.Context, payload dto.RegisterUserSchema) (interface{}, int, error)
	VerificationUser(ctx context.Context, payload dto.VerificationUser) (string, int, error)
	ResendVerificationUser(ctx context.Context, payload dto.ResendVerificationUser) (interface{}, int, error)
	ForgotPassword(ctx context.Context, email string, activitySource *models.UsersActivityHistory) (*ForgotPasswordResponse, int, error)
	ResendForgotPassword(ctx context.Context, payload dto.ResendForgotPassword) (interface{}, int, error)
	ResetPassword(ctx context.Context, payload dto.ResetPassword, activitySource *models.UsersActivityHistory) (int, error)

	// Admin Section
	ValidateUserAdmin(ctx context.Context, payload dto.ValidateUserSchema) (*UserTokenResponse, int, error)
}

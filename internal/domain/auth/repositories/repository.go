package repositories

import (
	"context"

	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/dtos"
	"gorm.io/gorm"
)

type repositoryAuth struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *repositoryAuth {
	return &repositoryAuth{db}
}

type AuthRepository interface {
	Register(ctx context.Context, payload *dto.RegisterUserSchema) (interface{}, error)
	// Verification(ctx context.Context, id string) (*model.User, error)
	// ResendVerification(ctx context.Context, email string) (*model.User, error)
	// ForgotPassword(ctx context.Context, payload *model.User) (*model.User, error)
	// ResendForgotPassword(ctx context.Context, payload *model.User) (*model.User, error)
}

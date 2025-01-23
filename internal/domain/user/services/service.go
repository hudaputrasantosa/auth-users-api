package services

import (
	"context"

	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/dtos"
	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	repo "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/repositories"
)

type serviceUser struct {
	userRepository repo.UserRepository
}

// User Service berfungsi untuk mengimplementasikan schema header yang di inisiasi di header service
func NewUserService(userRepository repo.UserRepository) *serviceUser {
	return &serviceUser{
		userRepository,
	}
}

// Interface User Service untuk mengetahui beberapa schema header yang tersedia pada User service
type UserService interface {
	Finds(ctx context.Context) (*[]model.User, error)
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Save(ctx context.Context, payload *dto.CreateUserSchema) (*model.User, error)
	Update(ctx context.Context, payload *dto.UpdateUserSchema) (*model.User, error)
	Delete(ctx context.Context, id string) error
}

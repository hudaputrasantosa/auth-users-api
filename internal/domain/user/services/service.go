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
	Finds(ctx context.Context) (*[]model.User, int, error)
	FindByID(ctx context.Context, id string) (*model.User, int, error)
	FindByEmail(ctx context.Context, email string) (*model.User, int, error)
	Save(ctx context.Context, payload *dto.CreateUserSchema) (*model.User, int, error)
	Update(ctx context.Context, payload *dto.UpdateUserSchema) (*model.User, int, error)
	Delete(ctx context.Context, id string) (int, error)
}

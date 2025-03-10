package repositories

import (
	"context"

	"github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	"gorm.io/gorm"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db}
}

type UserRepository interface {
	// User Interface
	Finds(ctx context.Context) (*[]models.User, error)
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Save(ctx context.Context, payload *models.User) (*models.User, error)
	Update(ctx context.Context, payload *models.User) (*models.User, error)
	UpdateStatusById(ctx context.Context, id string) (*models.User, error)
	Delete(ctx context.Context, user models.User) error

	// User Activity History Interface
	FindByUser(ctx context.Context, userId string) (*[]models.UsersActivityHistory, error)
	SaveActivity(ctx context.Context, payload *models.UsersActivityHistory) (*models.UsersActivityHistory, error)
}

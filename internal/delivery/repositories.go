package delivery

import (
	authRepo "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/repositories"
	userRepo "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository userRepo.UserRepository
	AuthRepository authRepo.AuthRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository: userRepo.NewUserRepository(db),
		AuthRepository: authRepo.NewAuthRepository(db),
	}
}

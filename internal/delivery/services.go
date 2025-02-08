package delivery

import (
	authService "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/services"
	userService "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/services"
	"github.com/redis/go-redis/v9"
)

type Services struct {
	UserService userService.UserService
	AuthService authService.AuthService
}

type Inject struct {
	Repository *Repositories
	Redis      *redis.Client
}

func NewService(inject Inject) *Services {
	return &Services{
		UserService: userService.NewUserService(inject.Repository.UserRepository),
		AuthService: authService.NewAuthService(inject.Repository.AuthRepository, inject.Repository.UserRepository, inject.Redis),
	}
}

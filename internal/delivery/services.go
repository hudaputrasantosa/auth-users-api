package delivery

import (
	authService "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/services"
	userService "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/services"
)

type Services struct {
	UserService userService.UserService
	AuthService authService.AuthService
}

type Inject struct {
	Repository *Repositories
}

func NewService(inject Inject) *Services {
	return &Services{
		UserService: userService.NewUserService(inject.Repository.UserRepository),
		AuthService: authService.NewAuthService(inject.Repository.AuthRepository, inject.Repository.UserRepository),
	}
}

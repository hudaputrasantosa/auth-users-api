package dto

import (
	"github.com/hudaputrasantosa/auth-users-api/internal/domain/entity"
)

type Users struct {
	Users []entity.User `json:"users"`
}

type CreateUserSchema struct {
    Name  	  string `json:"name" validate:"required"`
    Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Phone     string   `json:"phone" validate:"required,number,min=12"`
	IsActive  bool   `json:"is_active,omitempty"`
}

type UpdateUserSchema struct {
    Name  string `json:"name,omitempty"`
    Username  string `json:"username,omitempty"`
	Phone     string   `json:"phone,omitempty" validate:"number,min=12"`
	IsActive  bool   `json:"is_active,omitempty"`
}

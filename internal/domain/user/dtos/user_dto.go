package dtos

import (
	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
)

type Users struct {
	Users []model.User `json:"users"`
}

type CreateUserSchema struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required,number,min=12"`
	IsActive bool   `json:"is_active,omitempty"`
}

type UpdateUserSchema struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Phone    string `json:"phone,omitempty" validate:"number,min=12"`
	IsActive bool   `json:"is_active,omitempty"`
}

type UpdateUserMemberSchema struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty" validate:"number,min=12"`
}

type UpdatePasswordMember struct {
	LastPassword string `json:"last_password" validate:"required"`
	NewPassword  string `json:"new_password" validate:"required"`
}

type DeactivatedAccount struct {
	Password string `json:"password" validate:"required"`
}

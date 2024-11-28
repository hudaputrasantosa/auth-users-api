package dto

type RegisterUserSchema struct {
    Name  	  string `json:"name" validate:"required"`
    Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Phone     string   `json:"phone" validate:"required"`
}

type ValidateUserSchema struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

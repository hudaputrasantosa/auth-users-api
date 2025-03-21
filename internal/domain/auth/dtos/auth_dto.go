package dtos

type RegisterUserSchema struct {
	Name     string `json:"name" validate:"required,min=4"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Phone    string `json:"phone" validate:"required"`
}

type ValidateUserSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type VerificationUser struct {
	Token string `json:"token" validate:"required"`
	Otp   string `json:"otp" validate:"required,min=6"`
}
type ResendVerificationUser struct {
	Token string `json:"token" validate:"required"`
}
type RequestForgotPassword struct {
	Email string `json:"email" validate:"required,email"`
}

type ResendForgotPassword struct {
	Token string `json:"token" validate:"required"`
}
type ResetPassword struct {
	Token       string `json:"token" validate:"required"`
	Otp         string `json:"otp" validate:"required,min=6"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

package rest

import (
	"github.com/gofiber/fiber/v2"

	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/dtos"
	authService "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/services"

	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/response"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/validation"
)

type handleAuth struct {
	authClient authService.AuthService
}

// Router Auth
func NewHandleAuthRoute(
	authClient authService.AuthService,
	router *fiber.App) {
	handlerAuth := handleAuth{
		authClient,
	}

	api := router.Group("/api")
	routerV1 := api.Group("/v1/auth")

	//Routes Version 1.0
	routerV1.Post("/register", handlerAuth.registerUser)
	routerV1.Post("/login", handlerAuth.validateUser)
	routerV1.Post("/verification", handlerAuth.verificationUser)
	routerV1.Post("/verification/resend", handlerAuth.resendVerificationUser)
	routerV1.Post("/forgot-password", handlerAuth.forgotPassword)
	routerV1.Post("/forgot-password/resend", handlerAuth.resendForgotPassword)
	routerV1.Post("/reset-password", handlerAuth.resetPassword)
}

// Handler / Controller Auth Service
func (h *handleAuth) validateUser(c *fiber.Ctx) error {
	ctx := c.Context()

	// Create or initial login struct payload
	var payload dto.ValidateUserSchema

	// Check received data from JSON body.
	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, "Failed parsing", err)
	}

	// Validate data before proceessing.
	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
	}

	res, status, err := h.authClient.ValidateUser(ctx, payload)
	if err != nil {
		if res != nil {
			return response.ErrorMessage(c, status, "Failed login", err, res)
		}
		return response.ErrorMessage(c, status, "Failed login", err)
	}

	return response.SuccessMessageWithData(c, status, "Success login", res)
}

func (h *handleAuth) registerUser(c *fiber.Ctx) error {
	ctx := c.Context()

	// Create or initial user struct payload
	var payload dto.RegisterUserSchema

	// Check received data from JSON body.
	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, "Failed parsing", err)
	}

	// check input validation
	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
	}

	res, status, err := h.authClient.RegisterUser(ctx, payload)
	if err != nil {
		return response.ErrorMessage(c, status, "Failed register", err)
	}

	// return success with token otp
	return response.SuccessMessageWithData(c, status, "Success Register", res)
}

func (h *handleAuth) verificationUser(c *fiber.Ctx) error {
	ctx := c.Context()
	// Create or initial user struct payload
	var payload dto.VerificationUser

	// Check received data from JSON body.
	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, "Failed parsing", err)
	}

	// check input validation
	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
	}

	res, status, err := h.authClient.VerificationUser(ctx, payload)
	if err != nil {
		return response.ErrorMessage(c, status, "Failed verification", err)
	}

	// return success with token otp
	return response.SuccessMessageWithData(c, status, "Verification user successfully", res)
}

func (h *handleAuth) resendVerificationUser(c *fiber.Ctx) error {
	ctx := c.Context()
	// Create or initial user struct payload
	var payload dto.ResendVerificationUser

	// Check received data from JSON body.
	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, "Failed parsing", err)
	}

	// check input validation
	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
	}

	res, status, err := h.authClient.ResendVerificationUser(ctx, payload)
	if err != nil {
		return response.ErrorMessage(c, status, "Failed verification", err)
	}

	// return success with token otp
	return response.SuccessMessageWithData(c, status, "Verification user successfully", res)
}

func (h *handleAuth) forgotPassword(c *fiber.Ctx) error {
	ctx := c.Context()
	// Create or initial user struct payload
	var payload dto.RequestForgotPassword

	// Check received data from JSON body.
	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, "Failed parsing", err)
	}

	// check input validation
	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
	}

	res, status, err := h.authClient.ForgotPassword(ctx, payload.Email)
	if err != nil {
		return response.ErrorMessage(c, status, "Failed Forgot Password", err)
	}
	return response.SuccessMessageWithData(c, status, "Request Forgot Password successfully", res)
}

func (h *handleAuth) resendForgotPassword(c *fiber.Ctx) error {
	ctx := c.Context()
	// Create or initial user struct payload
	var payload dto.ResendForgotPassword

	// Check received data from JSON body.
	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, "Failed parsing", err)
	}

	// check input validation
	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
	}

	res, status, err := h.authClient.ResendForgotPassword(ctx, payload)
	if err != nil {
		return response.ErrorMessage(c, status, "Failed resend forgot password", err)
	}

	// return success with token otp
	return response.SuccessMessageWithData(c, status, "resend forgot password successfully", res)
}

func (h *handleAuth) resetPassword(c *fiber.Ctx) error {
	ctx := c.Context()
	// Create or initial user struct payload
	var payload dto.ResetPassword

	// Check received data from JSON body.
	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, "Failed parsing", err)
	}

	// check input validation
	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
	}

	status, err := h.authClient.ResetPassword(ctx, payload)
	if err != nil {
		return response.ErrorMessage(c, status, "Failed Reset Password", err)
	}
	return response.SuccessMessageWithData(c, status, "Reset Password successfully", "")
}

package rest

import (
	"github.com/gofiber/fiber/v2"

	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/dtos"
	authService "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/services"

	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/response"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/validation"
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

	res, err := h.authClient.ValidateUser(ctx, payload)
	if err != nil {
		return response.ErrorMessage(c, fiber.StatusInternalServerError, "Failed", err)
	}

	return response.SuccessMessageWithData(c, fiber.StatusOK, "Success login", res)
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

	res, err := h.authClient.RegisterUser(ctx, payload)
	if err != nil {
		return response.ErrorMessage(c, fiber.StatusInternalServerError, "Failed to create user", err)
	}

	// return success with token otp
	return response.SuccessMessageWithData(c, fiber.StatusOK, "Success Register", res)
}

package rest

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth/gothic"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/hudaputrasantosa/auth-users-api/internal/config"
	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/dtos"
	authService "github.com/hudaputrasantosa/auth-users-api/internal/domain/auth/services"

	"github.com/hudaputrasantosa/auth-users-api/pkg/middleware"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils"
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
	// Auth Member Router Version 1.0
	memberRouterV1 := api.Group("/v1/member/auth")
	memberRouterV1.Post("/register", middleware.RateLimit(3, 15, nil), handlerAuth.registerUser)
	memberRouterV1.Post("/login", middleware.RateLimit(5, 15, nil), handlerAuth.validateUser)
	memberRouterV1.Get("/:provider", middleware.RateLimit(5, 15, nil), handlerAuth.googleAuthentication)
	memberRouterV1.Post("/google/v2", middleware.RateLimit(5, 15, nil), handlerAuth.googleAuthenticationV2)
	memberRouterV1.Post("/:provider/callback", middleware.RateLimit(5, 15, nil), handlerAuth.googleAuthenticationCallback)
	memberRouterV1.Post("/google/callback/v2", middleware.RateLimit(5, 15, nil), handlerAuth.googleAuthenticationCallbackV2)
	memberRouterV1.Post("/verification", middleware.RateLimit(3, 15, nil), handlerAuth.verificationUser)
	memberRouterV1.Post("/verification/resend", middleware.RateLimit(1, 60, nil), handlerAuth.resendVerificationUser)
	memberRouterV1.Post("/forgot-password", middleware.RateLimit(3, 15, nil), handlerAuth.forgotPassword)
	memberRouterV1.Post("/forgot-password/resend", middleware.RateLimit(1, 60, nil), handlerAuth.resendForgotPassword)
	memberRouterV1.Post("/reset-password", middleware.RateLimit(3, 15, nil), handlerAuth.resetPassword)

	// Auth Admin Router Version 1.0
	adminRouterV1 := api.Group("/v1/admin/auth")
	adminRouterV1.Post("/login", middleware.RateLimit(5, 15, nil), handlerAuth.validateUserAdmin)
}

// Handler / Controller Auth Service
func (h *handleAuth) googleAuthentication(c *fiber.Ctx) error {
	provider := c.Params("provider")
	fmt.Println("Auth provider:", provider)

	// Konversi fiber.Ctx ke http.ResponseWriter dan *http.Request
	req := new(http.Request)
	fasthttpadaptor.ConvertRequest(c.Context(), req, true)
	res := &utils.FiberResponseWriter{Ctx: c}

	// Redirect ke halaman login Google
	gothic.BeginAuthHandler(res, req)
	return nil
}

func (h *handleAuth) googleAuthenticationV2(c *fiber.Ctx) error {
	path := config.ConfigGoogle()
	url := path.AuthCodeURL("state")
	return c.Redirect(url)
}

func (h *handleAuth) googleAuthenticationCallback(c *fiber.Ctx) error {
	// Konversi fiber.Ctx ke http.ResponseWriter dan *http.Request
	req := new(http.Request)
	fasthttpadaptor.ConvertRequest(c.Context(), req, true)
	res := &utils.FiberResponseWriter{Ctx: c}

	// Ambil user dari callback OAuth
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		return response.ErrorMessage(c, fiber.StatusUnauthorized, "Failed login", err)
	}

	return response.SuccessMessageWithData(c, fiber.StatusOK, "Success redirect login", user)
}

func (h *handleAuth) googleAuthenticationCallbackV2(c *fiber.Ctx) error {
	token, error := config.ConfigGoogle().Exchange(c.Context(), c.FormValue("code"))
	if error != nil {
		return response.ErrorMessage(c, fiber.StatusInternalServerError, "Failed login", error)
	}

	user := config.GetUserInfo(token.AccessToken)

	return response.SuccessMessageWithData(c, fiber.StatusOK, "Success login", user.Email)
}

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

	activitySource := utils.DetectDevice(c)
	res, status, err := h.authClient.ValidateUser(ctx, payload, &activitySource)
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
	activitySource := utils.DetectDevice(c)
	res, status, err := h.authClient.ForgotPassword(ctx, payload.Email, &activitySource)
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
	activitySource := utils.DetectDevice(c)
	status, err := h.authClient.ResetPassword(ctx, payload, &activitySource)
	if err != nil {
		return response.ErrorMessage(c, status, "Failed Reset Password", err)
	}
	return response.SuccessMessageWithData(c, status, "Reset Password successfully", "")
}

func (h *handleAuth) validateUserAdmin(c *fiber.Ctx) error {
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

	res, status, err := h.authClient.ValidateUserAdmin(ctx, payload)
	if err != nil {
		if res != nil {
			return response.ErrorMessage(c, status, "Failed login", err, res)
		}
		return response.ErrorMessage(c, status, "Failed login", err)
	}

	return response.SuccessMessageWithData(c, status, "Success login", res)
}

package rest

import (
	"github.com/gofiber/fiber/v2"
	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/dtos"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/response"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/validation"
)

func (h *handleUser) getUserMemberById(c *fiber.Ctx) error {
	ctx := c.Context()
	userId := c.Params("id")

	res, status, err := h.userClient.FindByID(ctx, userId)
	if err != nil {
		return response.ErrorMessage(c, status, "User not found", err)
	}

	return response.SuccessMessageWithData(c, status, "Success get users data", res)
}

func (h *handleUser) updateProfile(c *fiber.Ctx) error {
	ctx := c.Context()
	userId := c.Params("id")
	var payload *dto.UpdateUserMemberSchema

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, err.Error(), err)
	}

	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
	}

	res, status, err := h.userClient.UpdateMemberById(ctx, userId, payload)
	if err != nil {
		return response.ErrorMessage(c, status, "Failed create user", err)
	}

	return response.SuccessMessageWithData(c, status, "Success created", res)
}

func (h *handleUser) updatePassword(c *fiber.Ctx) error {
	ctx := c.Context()
	userId := c.Params("id")
	var payload *dto.UpdatePasswordMember

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, err.Error(), err)
	}

	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
	}

	status, err := h.userClient.UpdatePassword(ctx, userId, payload)
	if err != nil {
		return response.ErrorMessage(c, status, "Failed create user", err)
	}

	return response.SuccessMessageWithData(c, status, "Success update password", nil)
}

func (h *handleUser) deactivatedAccount(c *fiber.Ctx) error {
	ctx := c.Context()
	userId := c.Params("id")
	var payload *dto.DeactivatedAccount

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, err.Error(), err)
	}

	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
	}

	status, err := h.userClient.DeactivatedAccount(ctx, userId, payload.Password)
	if err != nil {
		return response.ErrorMessage(c, status, "Failed deactivated user", err)
	}

	return response.SuccessMessageWithData(c, status, "Success deactivated account", nil)
}

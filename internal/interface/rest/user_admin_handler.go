package rest

import (
	"github.com/gofiber/fiber/v2"
	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/dtos"
	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/response"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/validation"
	"github.com/morkid/paginate"
)

func (h *handleUser) createUser(c *fiber.Ctx) error {
	ctx := c.Context()

	var payload *dto.CreateUserSchema

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, err.Error(), err)
	}

	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
	}

	res, status, err := h.userClient.Save(ctx, payload)
	if err != nil {
		return response.ErrorMessage(c, status, "Failed create user", err)
	}

	return response.SuccessMessageWithData(c, status, "Success created", res)
}

func (h *handleUser) findUsersPagination(c *fiber.Ctx) error {
	db := database.DB.Db

	var users []model.User
	pg := paginate.New()
	stmt := db.Find(&users)
	page := pg.With(stmt).Request(c.Request()).Response(&users)

	return c.Status(200).JSON(page)
}

func (h *handleUser) findUsers(c *fiber.Ctx) error {
	ctx := c.Context()

	res, status, err := h.userClient.Finds(ctx)
	if err != nil {
		return response.ErrorMessage(c, status, "Data not found", nil)
	}

	return response.SuccessMessageWithData(c, fiber.StatusOK, "Success get users data", res)
}

func (h *handleUser) findUserById(c *fiber.Ctx) error {
	ctx := c.Context()
	userId := c.Params("id")

	res, status, err := h.userClient.FindByID(ctx, userId)
	if err != nil {
		return response.ErrorMessage(c, status, "Data not found", err)
	}

	return response.SuccessMessageWithData(c, status, "Success get users data", res)
}

func (h *handleUser) deleteUserById(c *fiber.Ctx) error {
	ctx := c.Context()
	userId := c.Params("id")

	status, err := h.userClient.Delete(ctx, userId)
	if err != nil {
		return response.ErrorMessage(c, status, "Failed delete", err)
	}

	return response.SuccessMessage(c, status, "Success delete")
}

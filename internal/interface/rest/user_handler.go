package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"

	dto "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/dtos"
	model "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	service "github.com/hudaputrasantosa/auth-users-api/internal/domain/user/services"
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database"
	"github.com/hudaputrasantosa/auth-users-api/internal/middleware"
	middle "github.com/hudaputrasantosa/auth-users-api/pkg/middleware"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/response"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/validation"
)

type handleUser struct {
	userClient service.UserService
}

// Router Users
func NewHandleUserRoute(userClient service.UserService, router *fiber.App) {
	handlerUser := handleUser{
		userClient,
	}

	api := router.Group("/api")

	// Router Group Version 1.0 and Middleware Group
	routerV1 := api.Group("/v1/users")
	routerV1.Use(middleware.Authorized(), middle.RateLimit(30, 60, nil))

	routerV1.Use(middleware.IsAdmin(handlerUser.userClient))
	routerV1.Post("/", handlerUser.createUser)
	routerV1.Get("/", handlerUser.findUsers)
	routerV1.Get("/pagination", handlerUser.findUsersPagination)
	routerV1.Get("/:id", handlerUser.findUserById)
	routerV1.Delete("/:id", handlerUser.deleteUserById)
}

// Handler / Controller User Service
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

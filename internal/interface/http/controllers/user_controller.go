package controllers

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/google/uuid"
	"github.com/morkid/paginate"

	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database"
	"github.com/hudaputrasantosa/auth-users-api/internal/interface/http/dto"
	"github.com/hudaputrasantosa/auth-users-api/internal/domain/entity"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/response"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/hash"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/validation"

)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	var payload *dto.CreateUserSchema;

	if err := c.BodyParser(&payload); err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, err.Error(), err)
	}

	if err := validation.ValidateStructDetail(payload); err != nil {
		return response.ErrorValidationMessage(c, fiber.StatusBadRequest, err)
	}

	now := time.Now();
	hashPassword, _ := hash.HashPassword(payload.Password)
	newUser := entity.User{
		ID:        uuid.New(),
        Name:      payload.Name,
        Username:  payload.Username,
        Email:     payload.Email,
        Password:  hashPassword,
        Phone:     payload.Phone,
        CreatedAt: now,
        UpdatedAt: now,
	}

	err := db.Create(&newUser).Error
	if err != nil {
		return response.ErrorMessage(c, fiber.StatusBadRequest, "Failed create user input", err)
	}
	return response.SuccessMessageWithData(c, fiber.StatusOK, "Success created", newUser)
}

func FindUsersPagination(c *fiber.Ctx) error {
	db := database.DB.Db

	var users []entity.User
	pg := paginate.New()
	stmt := db.Find(&users)
	page := pg.With(stmt).Request(c.Request()).Response(&users)

	return c.Status(200).JSON(page)
}

func FindUsers(c *fiber.Ctx) error {
	db := database.DB.Db

	var users []entity.User
	db.Find(&users)

	if len(users) == 0 {
		return response.ErrorMessage(c, fiber.StatusNotFound, "User data notfound", nil)
	}

	return response.SuccessMessageWithData(c, fiber.StatusOK, "Success get users data", users)
}

func FindUserById(c *fiber.Ctx) error {
	db := database.DB.Db
	userId := c.Params("id")
	var user *entity.User

	result := db.First(&user, "id = ?", userId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.ErrorMessage(c, fiber.StatusNotFound, "User id not found", nil)
		}
			return response.ErrorMessage(c, fiber.StatusInternalServerError, err.Error(), err)
	}

	return response.SuccessMessageWithData(c, fiber.StatusOK, "Success get users data", user)
}

func DeleteUserById(c *fiber.Ctx) error {
	db := database.DB.Db
	userId := c.Params("id")
	var user entity.User

	result := db.First(&user, "id = ?", userId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.ErrorMessage(c, fiber.StatusNotFound, "User id not found", nil)
		}
			return response.ErrorMessage(c, fiber.StatusInternalServerError, err.Error(), err)
	}

	db.Delete(&user, "id = ?", userId)
	return response.SuccessMessage(c, fiber.StatusOK, "Success delete users data")
}

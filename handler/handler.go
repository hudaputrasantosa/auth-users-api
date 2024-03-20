package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hudaputrasantosa/auth-users-api/database"
	"github.com/hudaputrasantosa/auth-users-api/model"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "err",
			"message": "error user input",
			"data":    err,
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "err",
			"message": "failed create user input",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Success created",
		"data":    user,
	})
}

func GetAllUser(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "Notfound",
			"message": "user data notfound",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "OK",
		"message": "success get data user",
		"data":    users,
	})

}

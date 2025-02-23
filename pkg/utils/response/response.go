package response

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/validation"
)

type Response struct {
	Error        bool        `json:"error,omitempty"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	ErrorCode    int         `json:"error_code,omitempty"`
	ErrorType    string      `json:"error_type,omitempty"`
	ErrorDetails string      `json:"error_details,omitempty"`
}

func SuccessMessage(c *fiber.Ctx, statusCode int, message string) error {
	response := &Response{
		Error:   false,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

func SuccessMessageWithData(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := &Response{
		Error:   false,
		Message: message,
		Data:    data,
	}
	return c.Status(statusCode).JSON(response)
}

func RespondWithPagination(c *fiber.Ctx, code int, message string, total int, page int, perPage int, dataName string, data interface{}) error {
	return c.Status(code).JSON(fiber.Map{
		"error":   false,
		"message": message,
		"data": fiber.Map{
			dataName:   data,
			"total":    total,
			"page":     page,
			"per_page": perPage,
		},
	})
}

func ErrorMessage(c *fiber.Ctx, statusCode int, message string, err error, data ...interface{}) error {
	response := &Response{
		Error:   true,
		Message: message,
	}

	if err != nil {
		response.ErrorDetails = fmt.Sprintf("%v", err)
	}

	if data != nil {
		response.Data = data
	}

	return c.Status(statusCode).JSON(response)
}

func ErrorValidationMessage(c *fiber.Ctx, statusCode int, dataValidation []*validation.ErrorResponse) error {
	response := &Response{
		Error:   true,
		Message: "Validation Error",
		Data:    dataValidation,
	}

	return c.Status(statusCode).JSON(response)
}

func ErrorMessageDetail(c *fiber.Ctx, statusCode int, errorCode int, errorType, message, details string, err error) error {
	response := &Response{
		Error:        true,
		Message:      message,
		ErrorCode:    errorCode,
		ErrorType:    errorType,
		ErrorDetails: fmt.Sprintf("%s : %v", details, err),
	}
	return c.Status(statusCode).JSON(response)
}

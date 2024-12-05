package validation

import (
	"sync"
	"strings"
	"regexp"


	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     sync.Once
)

func GetValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New(validator.WithRequiredStructEnabled())
	})

	return validate
}

func ValidateStruct(s interface{}) error {
	return GetValidator().Struct(s)
}

type ErrorResponse struct {
	FieldName  string `json:"fieldName"`
	Message    string `json:"message,omitempty"`
}

func ValidateStructDetail[T any](payload T) []*ErrorResponse {
	var errFields []*ErrorResponse
	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("phone", phoneValidation)
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errField ErrorResponse
			switch err.Tag() {
				case "required":
                    errField.FieldName = strings.ToLower(err.Field())
                    errField.Message = err.Field() + " field is required"
                case "email":
                    errField.FieldName = strings.ToLower(err.Field())
                    errField.Message = "Invalid email format"
				case "min":
					errField.FieldName = strings.ToLower(err.Field())
					errField.Message = err.Field() + " must be minimum " + err.Param() + " characters"
				case "max":
					errField.FieldName = strings.ToLower(err.Field())
					errField.Message = err.Field() + " maximum allowed is" + err.Param() + " characters"
				case "phone":
					errField.FieldName = strings.ToLower(err.Field())
					errField.Message = "Invalid Phone number format"
                default:
                    errField.FieldName = strings.ToLower(err.Field())
                    errField.Message = err.Tag()
			}
			errFields = append(errFields, &errField)
		}
	}

	if len(errFields) == 0 {
		return nil
	}

	return errFields
}

func phoneValidation(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	if len(phoneNumber) < 10 || len(phoneNumber) > 13 {
		return false
	}
	phoneRegex, _ := regexp.Compile(`^(0|\\+62|062|62)[0-9]+$`)
	return phoneRegex.MatchString(phoneNumber)
}

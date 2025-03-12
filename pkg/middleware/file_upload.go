package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hudaputrasantosa/auth-users-api/pkg/utils/response"
)

func ExtensionImageFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var allowedExtensions = []string{"jpg", "jpeg", "png"}
		file, err := c.FormFile("file")
		if err != nil {
			return response.ErrorMessage(c, fiber.StatusBadRequest, "file is required", nil)
		}

		// Get extension file
		fileName := file.Filename
		fileExt := strings.ToLower(fileName[strings.LastIndex(fileName, ".")+1:])

		// Check extension file allowed
		isAllowed := false
		for _, ext := range allowedExtensions {
			if fileExt == ext {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			return response.ErrorMessage(c, fiber.StatusBadRequest, "file extension not allowed", nil)
		}

		return c.Next()
	}
}

func FileSizeValidator(maxSize *int64) fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return response.ErrorMessage(c, fiber.StatusBadRequest, "file is required", nil)
		}

		if maxSize == nil {
			defaultSize := int64(1)
			maxSize = &defaultSize
		}

		// Check file size
		if file.Size > *maxSize {
			return response.ErrorMessage(c, fiber.StatusBadRequest, "file size exceeds the limit", nil)
		}

		return c.Next()
	}
}

package config

import (
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cldSecret := Config("CLOUDINARY_API_SECRET_KEY")
	cldName := Config("CLOUDINARY_CLOUD_NAME")
	cldKey := Config("CLOUDINARY_API_KEY")

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Cloudinary: %v", err)
	}

	return cld, nil
}

package config

import (
	"fmt"
	"os"

	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
	"github.com/joho/godotenv"
)

func Config(key string, fallback ...string) string {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error load env")
	}
	value := os.Getenv(key)
	if value == "" && len(fallback) > 0 {
		return fallback[0]
	}
	return value

}

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(name string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch name {
	case "postgres":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			Config("DB_HOST"),
			Config("DB_PORT"),
			Config("DB_USER"),
			Config("DB_PASSWORD"),
			Config("DB_NAME"),
			Config("DB_SSL_MODE"),
			Config("DB_TIMEZONE"),
		)
	case "mysql":
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			Config("DB_USER"),
			Config("DB_PASSWORD"),
			Config("DB_HOST"),
			Config("DB_PORT"),
			Config("DB_NAME"),
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			Config("REDIS_HOST"),
			Config("REDIS_PORT"),
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			Config("SERVER_HOST"),
			Config("SERVER_PORT"),
		)
	case "cloudinary":
		url = fmt.Sprintf(
			"cloudinary://%s:%s@%s",
			Config("CLOUDINARY_API_KEY"),
			Config("CLOUDINARY_API_SECRET_KEY"),
			Config("CLOUDINARY_CLOUD_NAME"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", name)
	}

	// Return connection URL.
	return url, nil
}

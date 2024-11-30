package connection

import (
	"fmt"

	"github.com/hudaputrasantosa/auth-users-api/internal/config"
)

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
			config.Config("DB_HOST"),
			config.Config("DB_PORT"),
			config.Config("DB_USER"),
			config.Config("DB_PASSWORD"),
			config.Config("DB_NAME"),
			config.Config("DB_SSL_MODE"),
			config.Config("DB_TIMEZONE"),
		)
	case "mysql":
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			config.Config("DB_USER"),
			config.Config("DB_PASSWORD"),
			config.Config("DB_HOST"),
			config.Config("DB_PORT"),
			config.Config("DB_NAME"),
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			config.Config("REDIS_HOST"),
			config.Config("REDIS_PORT"),
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s",
			// config.Config("SERVER_HOST"),
			config.Config("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", name)
	}

	// Return connection URL.
	return url, nil
}

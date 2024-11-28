package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"

)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error load env")
	}
	return os.Getenv(key)
}

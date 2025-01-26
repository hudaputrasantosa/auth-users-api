package migration

import (
	"fmt"
	// "gorm.io/gorm"

	"github.com/hudaputrasantosa/auth-users-api/internal/config"
	"github.com/hudaputrasantosa/auth-users-api/internal/domain/user/models"
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database"
	"github.com/hudaputrasantosa/auth-users-api/pkg/hash"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
)

func Migration() {
	err := database.DB.Db.AutoMigrate(&models.User{})
	if err != nil {
		logger.Fatal("Failed to migrate...")
	} else {
		logger.Info("Migrated successfully")

		password, _ := hash.HashPassword(config.Config("USER_PASSWORD_MIGRATION"))

		var user *models.User

		// check user email
		email := config.Config("USER_EMAIL_MIGRATION")
		if res := database.DB.Db.First(&user, "email = ?", email); res != nil {
			fmt.Printf("User with email %s already exists", user.Email)
		} else {
			user := &models.User{
				Name:     "Admin Migration",
				Username: "admin_migration",
				Email:    email,
				Password: password,
				Phone:    "6285156890287",
				Role:     "admin",
				IsActive: true,
			}
			if err := database.DB.Db.Create(&user).Error; err != nil {
				fmt.Printf("Failed to create user: %v\n", err)
			}

			fmt.Println("User created successfully")
		}

	}

}

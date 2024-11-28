package migration

import (
	"github.com/hudaputrasantosa/auth-users-api/internal/infrastructure/database"
	"github.com/hudaputrasantosa/auth-users-api/internal/domain/entity"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"

)

func Migration() {
	err := database.DB.Db.AutoMigrate(&entity.User{})
	if err != nil {
		logger.Fatal("Failed to migrate...")
	} else {
		logger.Info("Migrated successfully")
	}

}

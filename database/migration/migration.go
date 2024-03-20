package migration

import (
	"fmt"
	"log"

	"github.com/hudaputrasantosa/auth-users-api/database"
	"github.com/hudaputrasantosa/auth-users-api/model"
)

func Migration() {
	err := database.DB.Db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate...")
	}

	fmt.Println("Migrated successfully")
}

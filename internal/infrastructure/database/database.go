package database

import (
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"github.com/hudaputrasantosa/auth-users-api/internal/config"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() {
	dsn, _ := config.ConnectionURLBuilder("postgres")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})

	if err != nil {
		logger.Fatal("failed to connect db", zap.Error(err))
		os.Exit(2)
	}

	logger.Info("Success connect to database")
	gormLogger.Default.LogMode(gormLogger.Info)

	DB = Dbinstance{
		Db: db,
	}

}

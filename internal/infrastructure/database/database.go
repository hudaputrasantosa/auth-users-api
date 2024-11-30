package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"go.uber.org/zap"

	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"
	"github.com/hudaputrasantosa/auth-users-api/pkg/helper/connection"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() {
	dsn, _ := connection.ConnectionURLBuilder("postgres")

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

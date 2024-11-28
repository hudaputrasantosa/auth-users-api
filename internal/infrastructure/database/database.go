package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"go.uber.org/zap"


	"github.com/hudaputrasantosa/auth-users-api/internal/config"
	"github.com/hudaputrasantosa/auth-users-api/pkg/logger"

)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		logger.Error("error conv string")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})

	if err != nil {
		logger.Fatal("failed to connect db", zap.Error(err))
		os.Exit(2)
	}

	logger.Info("success to connect db")
	gormLogger.Default.LogMode(gormLogger.Info)

	DB = Dbinstance{
		Db: db,
	}

}

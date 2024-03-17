package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hudaputrasantosa/auth-users-api/config"
	"github.com/hudaputrasantosa/auth-users-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("error conv string")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("failed to connect db", err)
		os.Exit(2)
	}

	log.Println("success to connect db")
	logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&model.User{})

	DB = Dbinstance{
		Db: db,
	}

}

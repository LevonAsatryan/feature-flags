package db

import (
	"fmt"
	"os"

	"github.com/LevonAsatryan/feature-flags/server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the db")
	}

	db.AutoMigrate(&models.Group{})

	return db
}

var DB *gorm.DB = Setup()

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
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the db")
	}

	db.AutoMigrate(
		&models.Group{},
		&models.FeatureFlag{},
	)

	return db
}

var DB *gorm.DB = Setup()

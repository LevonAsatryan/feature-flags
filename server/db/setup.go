package db

import (
	"fmt"
	"os"

	"github.com/LevonAsatryan/feature-flags/server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB;

func Setup() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to postgres")
	}

	// Check if the DB already exists

	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", os.Getenv("POSTGRES_NAME"))
	res := db.Exec(stmt)

	if res.Error != nil {
		createDBCommand := fmt.Sprintf("CREATE DATABASE %s;", os.Getenv("POSTGRES_NAME"))
		db.Exec(createDBCommand)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
	}
	

	dsn = fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_NAME"),
	)
	
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the db")
	}

	db.AutoMigrate(
		&models.Group{},
		&models.FeatureFlag{},
	)

	return db
}


package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup() {
	db, err := gorm.Open(postgres.Open("feature_flags"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to the db")
	}

}

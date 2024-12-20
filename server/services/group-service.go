package services

import (
	postgres "github.com/LevonAsatryan/feature-flags/server/db"
	"github.com/LevonAsatryan/feature-flags/server/models"
)

var db = postgres.DB

func GetGroups() ([]models.Group, error) {
	var groups []models.Group

	err := db.Find(&groups).Error
	return groups, err
}

func CreateGroup(group *models.Group) error {
	return db.Create(&group).Error
}

func UpdateGroup(group *models.Group) error {
	return db.Save(&group).Error
}

func DeleteGroup(group *models.Group) error {
	return db.Delete(group).Error
}

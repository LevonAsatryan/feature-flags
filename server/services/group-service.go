package services

import (
	"fmt"

	postgres "github.com/LevonAsatryan/feature-flags/server/db"
	"github.com/LevonAsatryan/feature-flags/server/models"
)

var db = postgres.DB

var RootGroupID string

/*
 * Create a root group, if none is found in the DB for default group
 */
func CheckRegisterRootGroup() error {
	rootGroup := models.Group{
		Name: "root",
	}

	fmt.Printf("group name: %v \n", rootGroup.Name)

	err := db.Where("name = ?", rootGroup.Name).Find(&rootGroup).Error

	if err != nil {
		return err
	}

	if rootGroup.ID == "" {
		err = db.Create(&rootGroup).Error
	}

	RootGroupID = rootGroup.ID

	return err
}

func GetGroups() ([]models.Group, error) {
	var groups []models.Group

	err := db.Find(&groups).Error
	return groups, err
}

func GetGroup(id string) (*models.Group, error) {
	group := &models.Group{}

	err := db.First(&group, "id = ?", id).Error

	return group, err
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

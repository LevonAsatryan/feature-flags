package services

import (
	"fmt"

	postgres "github.com/LevonAsatryan/feature-flags/server/db"
	"github.com/LevonAsatryan/feature-flags/server/models"
)


var RootGroupID string

/*
 * Create a root group, if none is found in the DB for default group
 */
func CheckRegisterRootGroup() error {
	rootGroup := models.Group{
		Name: "root",
	}

	fmt.Printf("group name: %v \n", rootGroup.Name)

	err := postgres.DB.Where("name = ?", rootGroup.Name).Find(&rootGroup).Error

	if err != nil {
		return err
	}

	if rootGroup.ID == "" {
		err = postgres.DB.Create(&rootGroup).Error
	}

	RootGroupID = rootGroup.ID

	return err
}

func GetGroups() ([]models.Group, error) {
	var groups []models.Group

	err := postgres.DB.Find(&groups).Error
	return groups, err
}

func GetGroup(id string) (*models.Group, error) {
	group := &models.Group{}

	err := postgres.DB.First(&group, "id = ?", id).Error

	return group, err
}

func CreateGroup(group *models.Group) error {
	return postgres.DB.Create(&group).Error
}

func UpdateGroup(group *models.Group) error {
	return postgres.DB.Save(&group).Error
}

func DeleteGroup(group *models.Group) error {
	return postgres.DB.Delete(group).Error
}

package services

import (
	"fmt"

	postgres "github.com/LevonAsatryan/feature-flags/server/db"
	"github.com/LevonAsatryan/feature-flags/server/models"
	"github.com/LevonAsatryan/feature-flags/server/repositories"
	"github.com/gin-gonic/gin"
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
	projections := []string{"id"}

	exsitedGroup, err := repositories.FindByName[models.Group](rootGroup.Name, projections)

	if err != nil {
		return err
	}

	if exsitedGroup.ID == "" {
		err = db.Create(&rootGroup).Error
	}

	RootGroupID = rootGroup.ID

	return err
}

func GetGroups() ([]models.Group, error) {
	var groups []models.Group

	groups, err := repositories.FindAll[models.Group]()
	return groups, err
}

func GetGroup(id string, ctx *gin.Context, projections []string) (models.Group, error) {
	return repositories.FindByID[models.Group](id, projections)
}

func CreateGroup(group *models.Group) error {
	return repositories.Create[models.Group](group)
}

func UpdateGroup(group *models.Group) error {
	return repositories.Update[models.Group](group)
}

func DeleteGroup(group *models.Group) error {
	return repositories.Delete[models.Group](group)
}

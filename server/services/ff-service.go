package services

import (
	"github.com/LevonAsatryan/feature-flags/server/models"
	postgres "github.com/LevonAsatryan/feature-flags/server/db"
)

func GetFFs() ([]models.FeatureFlag, error) {
	var ffs []models.FeatureFlag

	err := postgres.DB.Find(&ffs).Error

	return ffs, err
}

func CreateFF(ff *models.FeatureFlag) error {
	if ff.GroupId == "" {
		ff.GroupId = RootGroupID
	}
	return postgres.DB.Create(&ff).Error
}

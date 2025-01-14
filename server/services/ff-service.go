package services

import (
	"github.com/LevonAsatryan/feature-flags/server/models"
)

func GetFFs() ([]models.FeatureFlag, error) {
	var ffs []models.FeatureFlag

	err := db.Find(&ffs).Error

	return ffs, err
}

func CreateFF(ff *models.FeatureFlag) error {
	if ff.GroupId == "" {
		ff.GroupId = RootGroupID
	}
	return db.Create(&ff).Error
}

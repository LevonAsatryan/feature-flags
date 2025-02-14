package services

import (
	"github.com/LevonAsatryan/feature-flags/server/models"
	"github.com/LevonAsatryan/feature-flags/server/repositories"
)

func GetFFs() ([]models.FeatureFlag, error) {
	var ffs []models.FeatureFlag

	ffs, err := repositories.FindAll[models.FeatureFlag]()

	return ffs, err
}

func CreateFF(ff *models.FeatureFlag) error {
	if ff.GroupId == "" {
		ff.GroupId = RootGroupID
	}
	return repositories.Create[models.FeatureFlag](ff)
}

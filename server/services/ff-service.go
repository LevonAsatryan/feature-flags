package services

import (
	"github.com/LevonAsatryan/feature-flags/server/models"
)

func GetFFs() (map[string][]models.FeatureFlag, error) {
	ffsMap := make(map[string][]models.FeatureFlag)

	var ffs []models.FeatureFlag

	err := db.Find(&ffs).Error

	for _, ff := range ffs {
		if len(ffsMap[ff.GroupId]) != 0 {
			ffsMap[ff.GroupId] = append(ffsMap[ff.GroupId], ff)
		} else {
			ffsMap[ff.GroupId] = []models.FeatureFlag{ff}
		}
	}

	return ffsMap, err
}

func CreateFF(ff *models.FeatureFlag) error {
	if ff.GroupId == "" {
		ff.GroupId = RootGroupID
	}
	return db.Create(&ff).Error
}

func GetFFsByGroupID(groupID string) ([]models.FeatureFlag, error) {
	var ffs []models.FeatureFlag

	err := db.Where("group_id = ?", groupID).Find(&ffs).Error

	return ffs, err
}

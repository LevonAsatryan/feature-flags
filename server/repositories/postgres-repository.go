package repositories

import (
	postgres "github.com/LevonAsatryan/feature-flags/server/db"
	"github.com/LevonAsatryan/feature-flags/server/models"
)

type Entity interface {
	models.FeatureFlag | models.Group
}

func FindAll[T Entity]() ([]T, error) {
	var entities []T

	err := postgres.DB.Find(&entities).Error
	return entities, err
}

func FindByID[T Entity](id string, projections []string) (T, error) {
	var entity T

	err := postgres.DB.Select(projections).Where("id = ?", id).First(&entity).Error
	return entity, err
}

func FindByName[T Entity](name string, projections []string) (T, error) {
	var entity T

	err := postgres.DB.Select(projections).Where("name = ?", name).Find(&entity).Error
	return entity, err
}

func Create[T Entity](entity *T) error {
	return postgres.DB.Create(&entity).Error
}

func Update[T Entity](entity *T) error {
	return postgres.DB.Updates(&entity).Error
}

func Delete[T Entity](entity *T) error {
	return postgres.DB.Delete(&entity).Error
}

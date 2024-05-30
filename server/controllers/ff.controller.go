package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LevonAsatryan/feature-flags/db"
	"github.com/LevonAsatryan/feature-flags/types"
	"github.com/gin-gonic/gin"
)

type CreateFFBody struct {
	Name string `json:"name" binding:"required"`
}

type UpdateFFBody struct {
	Value bool `json:"value"`
}

type FFController struct {
	DB *db.DB
}

func (c *FFController) CreateFF(context *gin.Context) ([]db.FeatureFlag, *types.Error) {
	var ffs []db.FeatureFlag

	var rb CreateFFBody

	if err := context.ShouldBindJSON(&rb); err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("Request body validation"),
		}
	}

	ffIDs, err := c.DB.CreateFF(rb.Name)

	if err != nil {
		return nil, err
	}

	for _, ffID := range ffIDs {
		ff, err := c.DB.GetFFById(ffID)

		if err != nil {
			return nil, &types.Error{
				Code: http.StatusInternalServerError,
				Err:  fmt.Errorf("failed to fetch the created feature flag"),
			}
		}

		ffs = append(ffs, *ff)
	}

	return ffs, nil
}

func (c *FFController) GetAll(context *gin.Context) ([]db.FeatureFlag, *types.Error) {
	ffs, err := c.DB.GetFFAll()

	if err != nil {
		return nil, err
	}

	return ffs, nil
}

func (c *FFController) GetById(context *gin.Context) (*db.FeatureFlag, *types.Error) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid feature flag id"),
		}
	}

	ff, tErr := c.DB.GetFFById(id)

	if tErr != nil {
		return nil, tErr
	}

	return ff, nil
}

func (c *FFController) Update(context *gin.Context) (*db.FeatureFlag, *types.Error) {
	var rb UpdateFFBody

	if err := context.ShouldBindJSON(&rb); err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid body"),
		}
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid id provided"),
		}
	}

	ff, tErr := c.DB.UpdateFF(id, rb.Value)

	if tErr != nil {
		return nil, tErr
	}

	return ff, nil
}

func (c *FFController) Delete(context *gin.Context) *types.Error {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid id provided"),
		}
	}
	_, tErr := c.DB.GetFFById(id)

	if tErr != nil {
		return tErr
	}

	tErr = c.DB.DeleteFF(id)

	return tErr
}

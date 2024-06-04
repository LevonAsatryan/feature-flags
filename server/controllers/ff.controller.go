package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LevonAsatryan/feature-flags/db"
	"github.com/LevonAsatryan/feature-flags/types"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateFFBody struct {
	Name string `json:"name" binding:"required"`
}

type UpdateFFBody struct {
	Value bool `json:"value"`
}

type FFController struct {
	DB  *db.Queries
	Ctx context.Context
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

	envs, err := c.DB.GetEnvAll(context)

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to get envs"),
		}
	}

	for _, env := range envs {
		ff, err := c.DB.CreateFF(c.Ctx, db.CreateFFParams{
			Name: pgtype.Text{
				String: rb.Name,
				Valid:  rb.Name != "",
			},
			EnvID: pgtype.Int4{
				Int32: int32(env.ID),
				Valid: true,
			},
		})

		if err != nil {
			return nil, &types.Error{
				Code: http.StatusInternalServerError,
				Err:  fmt.Errorf("failed to create ff"),
			}
		}

		ffs = append(ffs, ff)
	}

	return ffs, nil
}

func (c *FFController) GetAll(context *gin.Context) ([]db.FeatureFlag, *types.Error) {
	ffs, err := c.DB.GetFeatureFlagAll(c.Ctx)

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to get ff"),
		}
	}

	return ffs, nil
}

func (c *FFController) GetByEnvId(context *gin.Context, envID int32) ([]db.FeatureFlag, *types.Error) {
	ffs, err := c.DB.GetFeatureFlagByEnvId(c.Ctx, pgtype.Int4{
		Int32: envID,
		Valid: true,
	})

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to get ff"),
		}
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

	ff, err := c.DB.GetFeatureFlag(c.Ctx, int32(id))

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to get ff by id %d", id),
		}
	}

	return &ff, nil
}

func (c *FFController) Update(context *gin.Context) *types.Error {
	var rb UpdateFFBody

	if err := context.ShouldBindJSON(&rb); err != nil {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid body"),
		}
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid id provided"),
		}
	}

	err = c.DB.UpdateFF(c.Ctx, db.UpdateFFParams{
		ID:    int32(id),
		Value: pgtype.Bool{Bool: rb.Value, Valid: true},
	})

	if err != nil {
		return &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to update ff"),
		}
	}

	return nil
}

func (c *FFController) Delete(context *gin.Context) *types.Error {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid id provided"),
		}
	}

	err = c.DB.DeleteFF(context, int32(id))

	if err != nil {
		return &types.Error{
			Code: http.StatusNotFound,
			Err:  fmt.Errorf("feature flag not found"),
		}
	}

	return nil
}

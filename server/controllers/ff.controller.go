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

type FFDTO struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Value   bool   `json:"value"`
	EnvID   int32  `json:"envID"`
	GroupID int32  `json:"groupID"`
}

func ffdtoFromFF(ff *db.FeatureFlag) *FFDTO {
	return &FFDTO{
		ID:      ff.ID,
		Name:    ff.Name.String,
		Value:   ff.Value.Bool,
		EnvID:   ff.EnvID.Int32,
		GroupID: ff.GroupID.Int32,
	}
}

func ffdtoFromFFArr(ffs []db.FeatureFlag) []FFDTO {
	ffdtos := make([]FFDTO, len(ffs))

	for i := range ffs {
		ffdtos[i] = *ffdtoFromFF(&ffs[i])
	}

	return ffdtos
}

type CreateFFBody struct {
	Name string `json:"name" binding:"required"`
}

type UpdateFFBody struct {
	Value bool `json:"value"`
}

type UpdateFFNameBody struct {
	NameFrom string `json:"nameFrom" binding:"required"`
	NameTo   string `json:"nameTo" binding:"required"`
}

type UpdateFFGroupBody struct {
	Name string `json:"name" binding:"required"`
}

type FFController struct {
	DB  *db.Queries
	Ctx context.Context
}

func (c *FFController) CreateFF(context *gin.Context) ([]*FFDTO, *types.Error) {
	var ffs []*FFDTO

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

	if rb.Name == "" {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("the name field can not be empty"),
		}
	}

	prevFFs, err := c.DB.GetFFByName(c.Ctx, pgtype.Text{
		String: rb.Name,
		Valid:  true,
	})

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to check the feature flags by name %s", rb.Name),
		}
	}

	if len(prevFFs) != 0 {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("ff with name '%s' already exists", rb.Name),
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

		ffs = append(ffs, ffdtoFromFF(&ff))
	}

	return ffs, nil
}

func (c *FFController) GetAll(context *gin.Context) ([]FFDTO, *types.Error) {
	ffs, err := c.DB.GetFFAll(c.Ctx)

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to get ff"),
		}
	}

	ffdtos := ffdtoFromFFArr(ffs)

	return ffdtos, nil
}

func (c *FFController) GetByEnvId(context *gin.Context, envID int32) ([]FFDTO, *types.Error) {
	ffs, err := c.DB.GetFFByEnvId(c.Ctx, pgtype.Int4{
		Int32: envID,
		Valid: true,
	})

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to get ff"),
		}
	}

	return ffdtoFromFFArr(ffs), nil
}

func (c *FFController) GetById(context *gin.Context) (*FFDTO, *types.Error) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid feature flag id"),
		}
	}

	ff, err := c.DB.GetFF(c.Ctx, int32(id))

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to get ff by id %d", id),
		}
	}

	return ffdtoFromFF(&ff), nil
}

func (c *FFController) Update(context *gin.Context) (*FFDTO, *types.Error) {
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

	err = c.DB.UpdateFF(c.Ctx, db.UpdateFFParams{
		ID:    int32(id),
		Value: pgtype.Bool{Bool: rb.Value, Valid: true},
	})

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to update ff"),
		}
	}

	ff, err := c.DB.GetFF(c.Ctx, int32(id))

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to fetch ff by id=%d", id),
		}
	}

	return ffdtoFromFF(&ff), nil
}

func (c *FFController) AddFFToGroup(context *gin.Context, groupID int32) *types.Error {
	var rb UpdateFFGroupBody

	if err := context.ShouldBindJSON(&rb); err != nil {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("Request body validation"),
		}
	}

	if rb.Name == "" {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("name of ff can not be empty"),
		}
	}

	err := c.DB.UpdateFFGroup(c.Ctx, db.UpdateFFGroupParams{
		Name: pgtype.Text{
			String: rb.Name,
			Valid:  true,
		},
		GroupID: pgtype.Int4{
			Int32: groupID,
			Valid: true,
		},
	})

	if err != nil {
		return &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to update ffs"),
		}
	}

	return nil
}

func (c *FFController) UpdateName(context *gin.Context) ([]FFDTO, *types.Error) {
	var rb UpdateFFNameBody

	if err := context.ShouldBindJSON(&rb); err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid body"),
		}
	}

	if rb.NameFrom == "" || rb.NameTo == "" {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("nameFrom and nameTo are required and can not be empty"),
		}
	}

	ffs, err := c.DB.GetFFByName(c.Ctx, pgtype.Text{
		String: rb.NameTo,
		Valid:  true,
	})

	if len(ffs) != 0 {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("feature flag with name '%s' already exists", rb.NameTo),
		}
	}

	ffs, err = c.DB.GetFFByName(c.Ctx, pgtype.Text{
		String: rb.NameFrom,
		Valid:  true,
	})

	if len(ffs) == 0 {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("feature flag with name '%s' does not exists", rb.NameFrom),
		}
	}

	err = c.DB.UpdateFFName(c.Ctx, db.UpdateFFNameParams{
		Name: pgtype.Text{
			String: rb.NameFrom,
			Valid:  true,
		},
		Name_2: pgtype.Text{
			String: rb.NameTo,
			Valid:  true,
		},
	})

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to update ff"),
		}
	}

	ffs, err = c.DB.GetFFByName(c.Ctx, pgtype.Text{
		String: rb.NameTo,
		Valid:  true,
	})

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to fetch feature flags with name %s", rb.NameTo),
		}
	}

	return ffdtoFromFFArr(ffs), nil
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

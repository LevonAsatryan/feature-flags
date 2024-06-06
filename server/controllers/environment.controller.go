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

type EnvController struct {
	DB  *db.Queries
	Ctx context.Context
}

type CreateEnvBody struct {
	Name       string `json:"name"`
	OriginHost string `json:"originHost"`
}

func (c *EnvController) CheckAndCreateEnv() *types.Error {
	envs, err := c.DB.GetEnvAll(c.Ctx)

	if err != nil {
		return &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("could not fetch environments"),
		}
	}

	if len(envs) == 0 {
		createEnvParams := db.CreateEnvParams{
			Name: pgtype.Text{
				Valid:  true,
				String: "dev",
			},
			OriginHost: pgtype.Text{
				Valid:  true,
				String: "localhost",
			},
		}
		c.DB.CreateEnv(c.Ctx, createEnvParams)
	}
	return nil
}

func (c *EnvController) GetEnvCount(context *gin.Context) (int64, *types.Error) {
	count, err := c.DB.GetEnvCount(c.Ctx)

	if err != nil {
		return 0, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to fetch the environment count"),
		}
	}

	return count, nil
}

func (c *EnvController) Create(context *gin.Context) (*db.Env, *types.Error) {
	var body CreateEnvBody

	if err := context.ShouldBindJSON(&body); err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("request body validation"),
		}
	}

	envName := body.Name
	originHost := body.OriginHost

	if envName == "" || originHost == "" {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("envName and originHost cannot be empty"),
		}
	}

	env, err := c.DB.CreateEnv(c.Ctx, db.CreateEnvParams{
		Name: pgtype.Text{
			String: envName,
			Valid:  true,
		},
		OriginHost: pgtype.Text{
			String: originHost,
			Valid:  true,
		},
	})

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("Could not create the env with name: %s and originHost: %s", envName, originHost),
		}
	}

	return &env, nil
}

func (c *EnvController) GetAll(context *gin.Context) ([]db.Env, *types.Error) {
	envs, err := c.DB.GetEnvAll(c.Ctx)

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	return envs, nil
}

func (c *EnvController) GetById(context *gin.Context) (*db.Env, *types.Error) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid id provided %s", context.Param("id")),
		}
	}

	env, err := c.DB.GetEnv(c.Ctx, int32(id))

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("env with id=%d not found", id),
		}
	}

	return &env, nil
}

func (c *EnvController) Delete(context *gin.Context) *types.Error {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid id provided %s", context.Param("id")),
		}
	}

	err = c.DB.DeleteEnv(context, int32(id))

	if err != nil {
		return &types.Error{
			Code: http.StatusInternalServerError,
			// Err:  fmt.Errorf("failed to delete environment by id=%d", id),
			Err: err,
		}
	}

	return nil
}

func (c *EnvController) Update(context *gin.Context) *types.Error {
	var body CreateEnvBody

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid id provided %s", context.Param("id")),
		}
	}

	_, err = c.DB.GetEnv(context, int32(id))

	if err != nil {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("failed to get the env by id=%d", id),
		}
	}

	if err := context.ShouldBindJSON(&body); err != nil {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("request body validation"),
		}
	}

	envName := body.Name
	originHost := body.OriginHost

	if envName == "" {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("environment name cannot be an empty string"),
		}
	}

	err = c.DB.UpdateEnv(context, db.UpdateEnvParams{
		ID:         int32(id),
		Name:       pgtype.Text{String: envName, Valid: true},
		OriginHost: pgtype.Text{String: originHost, Valid: true},
	})

	if err != nil {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("failed to update environment"),
		}
	}
	return nil
}

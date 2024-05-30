package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LevonAsatryan/feature-flags/db"
	"github.com/LevonAsatryan/feature-flags/types"
	"github.com/gin-gonic/gin"
)

type EnvController struct {
	DB *db.DB
}

type CreateEnvBody struct {
	Name string `json:"name"`
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

	envs, err := c.DB.GetEnvAll()

	if err != nil {
		return nil, err
	}

	existingEnvId := envs[0].ID

	env, tErr := c.DB.CreateEnv(envName)

	if tErr != nil {
		return nil, tErr
	}

	tErr = c.DB.CopyFromExisting(existingEnvId, env.ID)

	if tErr != nil {
		return nil, tErr
	}

	return env, nil
}

func (c *EnvController) GetAll(context *gin.Context) ([]db.Env, *types.Error) {
	envs, err := c.DB.GetEnvAll()

	if err != nil {
		return nil, err
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

	env, tErr := c.DB.GetEnvById(id)

	if tErr != nil {
		return nil, tErr
	}

	return env, nil
}

func (c *EnvController) Delete(context *gin.Context) *types.Error {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid id provided %s", context.Param("id")),
		}
	}

	tErr := c.DB.DeleteFFsByEnvId(id)

	if tErr != nil {
		return tErr
	}

	tErr = c.DB.DeleteEnv(id)

	return tErr
}

func (c *EnvController) Update(context *gin.Context) (*db.Env, *types.Error) {
	var body CreateEnvBody

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid id provided %s", context.Param("id")),
		}
	}

	_, tErr := c.DB.GetEnvById(id)

	if tErr != nil {
		return nil, tErr
	}

	if err := context.ShouldBindJSON(&body); err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("request body validation"),
		}
	}

	envName := body.Name

	if envName == "" {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("environment name cannot be an empty string"),
		}
	}

	env, tErr := c.DB.UpdateEnv(id, envName)

	return env, tErr
}

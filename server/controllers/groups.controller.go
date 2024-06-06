package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/LevonAsatryan/feature-flags/db"
	"github.com/LevonAsatryan/feature-flags/types"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type GroupsController struct {
	DB  *db.Queries
	Ctx context.Context
}

type CreateGroupBody struct {
	Name string `json:"name" binding:"required"`
}

func (c *GroupsController) Create(ctx *gin.Context, envs []db.Env) ([]db.Group, *types.Error) {

	var rb CreateGroupBody
	var groups []db.Group

	if err := ctx.ShouldBindJSON(&rb); err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("Request body validation"),
		}
	}

	for _, env := range envs {
		group, err := c.DB.CreateGroup(c.Ctx, db.CreateGroupParams{
			Name:  pgtype.Text{String: rb.Name, Valid: true},
			EnvID: pgtype.Int4{Int32: int32(env.ID), Valid: true},
		})

		if err != nil {
			return nil, &types.Error{
				Code: http.StatusInternalServerError,
				Err:  fmt.Errorf("failed to create group with name %s for the enviornmentID %d", rb.Name, env.ID),
			}
		}

		groups = append(groups, group)
	}

	return groups, nil
}

func (c *GroupsController) GetAll(ctx *gin.Context) ([]db.Group, *types.Error) {
	groups, err := c.DB.GetGroupsAll(c.Ctx)

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to fetch groups"),
		}
	}

	if len(groups) == 0 {
		groups = make([]db.Group, 0)
	}

	return groups, nil
}

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

type GroupDTO struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	EnvID int32  `json:"envID"`
}

func groupDTOFromGroup(g *db.Group) *GroupDTO {
	return &GroupDTO{
		ID:    g.ID,
		Name:  g.Name.String,
		EnvID: g.ID,
	}
}

func groupDTOFromGroupArr(gs []db.Group) []GroupDTO {
	gdtos := make([]GroupDTO, len(gs))

	for i := range gs {
		gdtos[i] = *groupDTOFromGroup(&gs[i])
	}

	return gdtos
}

type GroupsController struct {
	DB  *db.Queries
	Ctx context.Context
}

type CreateGroupBody struct {
	Name string `json:"name" binding:"required"`
}

func (c *GroupsController) Create(ctx *gin.Context, envs []EnvDTO) ([]GroupDTO, *types.Error) {

	var rb CreateGroupBody
	var groups []db.Group

	if err := ctx.ShouldBindJSON(&rb); err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("request body validation"),
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

	return groupDTOFromGroupArr(groups), nil
}

func (c *GroupsController) GetAll(ctx *gin.Context) ([]GroupDTO, *types.Error) {
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

	return groupDTOFromGroupArr(groups), nil
}

func (c *GroupsController) GetById(ctx *gin.Context) (*GroupDTO, *types.Error) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid groups id"),
		}
	}

	group, err := c.DB.GetGroupById(c.Ctx, int32(id))

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusNotFound,
			Err:  fmt.Errorf("group with id %d not found", id),
		}
	}

	return groupDTOFromGroup(&group), nil
}

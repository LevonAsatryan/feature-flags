package controllers

import (
	"fmt"
	"net/http"

	"github.com/LevonAsatryan/feature-flags/server/middlewares"
	"github.com/LevonAsatryan/feature-flags/server/models"
	"github.com/LevonAsatryan/feature-flags/server/services"
	"github.com/gin-gonic/gin"
)

func RegisterGroupRoutes(r *gin.Engine) {
	api := r.Group("/groups")
	err := services.CheckRegisterRootGroup()

	if err != nil {
		panic(err)
	}

	api.GET("", func(ctx *gin.Context) {
		groups, err := services.GetGroups()

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Errorf("failed to fetch the groups")})
			return
		}

		ctx.JSON(http.StatusOK, groups)
	})

	api.GET("/:id", middlewares.ValidateId, func(ctx *gin.Context) {
		id := ctx.Param("id")
		projections := []string{"id", "Name", "CreatedAt", "UpdatedAt"}

		group, err := services.GetGroup(id, ctx, projections)

		if err != nil && err.Error() == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("group with id: %s not found", id)})
			return
		}

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, group)
	})

	api.POST("", func(ctx *gin.Context) {
		var group models.Group
		if err := ctx.ShouldBindJSON(&group); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if err := services.CreateGroup(&group); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": "Group created successfully"})
	})

	api.PUT("/:id", middlewares.ValidateId, func(ctx *gin.Context) {
		var group models.Group
		if err := ctx.ShouldBindJSON(&group); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		group.ID = ctx.Param("id")

		if err := services.UpdateGroup(&group); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Group updated successfully"})
	})

	api.DELETE("/:id", middlewares.ValidateId, func(ctx *gin.Context) {
		var group models.Group
		group.ID = ctx.Param("id")

		if err := services.DeleteGroup(&group); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Group deleted successfully"})
	})
}

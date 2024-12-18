package controllers

import (
	"net/http"

	"github.com/LevonAsatryan/feature-flags/server/middlewares"
	"github.com/LevonAsatryan/feature-flags/server/models"
	groupService "github.com/LevonAsatryan/feature-flags/server/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/groups")

	api.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, groupService.GetGroups())
	})

	api.POST("", func(ctx *gin.Context) {
		var group models.Group
		if err := ctx.ShouldBindJSON(&group); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := groupService.CreateGroup(&group); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": "Group created successfully"})
	})

	api.PUT("/:id", middlewares.ValidateId, func(ctx *gin.Context) {
		var group models.Group
		if err := ctx.ShouldBindJSON(&group); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		group.ID = ctx.Param("id")

		if err := groupService.UpdateGroup(&group); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Group updated successfully"})
	})

	api.DELETE("/:id", middlewares.ValidateId, func(ctx *gin.Context) {
		var group models.Group
		group.ID = ctx.Param("id")

		if err := groupService.DeleteGroup(&group); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Group deleted successfully"})
	})
}

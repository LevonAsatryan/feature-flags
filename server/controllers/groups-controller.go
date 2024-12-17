package controllers

import (
	"net/http"

	postgres "github.com/LevonAsatryan/feature-flags/server/db"
	"github.com/LevonAsatryan/feature-flags/server/models"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/groups")
	api.POST("", func(ctx *gin.Context) {
		var group models.Group
		if err := ctx.ShouldBindJSON(&group); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		postgres.DB.Create(&group)
		ctx.JSON(http.StatusCreated, gin.H{"message": "Group created successfully"})
	})
}

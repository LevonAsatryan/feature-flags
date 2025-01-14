package controllers

import (
	"fmt"
	"net/http"

	"github.com/LevonAsatryan/feature-flags/server/models"
	"github.com/LevonAsatryan/feature-flags/server/services"
	"github.com/gin-gonic/gin"
)

func RegisterFFRoutes(r *gin.Engine) {
	api := r.Group("/feature-flags")

	api.GET("", func(ctx *gin.Context) {
		ffs, err := services.GetFFs()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Errorf("Failed to fetch feature flags"))
			return
		}

		ctx.JSON(http.StatusOK, ffs)
	})

	api.POST("", func(ctx *gin.Context) {
		var ff models.FeatureFlag

		if err := ctx.ShouldBindJSON(&ff); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := services.CreateFF(&ff); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "FF created successfully"})
	})
}

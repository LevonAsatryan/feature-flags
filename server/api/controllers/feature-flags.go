package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeatureFlag struct {
	name    string
	enabled bool
}

func AddHandlers(router *gin.Engine) {
	ffs := router.Group("/ffs")
	{
		ffs.GET("/", handleGetFeatureFlags)
	}
}

func handleGetFeatureFlags(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

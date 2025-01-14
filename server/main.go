package main

import (
	"log"
	"net/http"
	"os"

	"github.com/LevonAsatryan/feature-flags/server/controllers"
	"github.com/LevonAsatryan/feature-flags/server/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.ErrorHandler())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	controllers.RegisterGroupRoutes(r)
	controllers.RegisterFFRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Run(":" + port)

	if err != nil {
		panic(err)
	}
}

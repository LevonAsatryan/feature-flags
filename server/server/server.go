package server

import (
	"github.com/LevonAsatryan/feature-flags/controllers"
	"github.com/LevonAsatryan/feature-flags/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	DB            *db.DB
	R             *gin.Engine
	FFController  controllers.FFController
	EnvController controllers.EnvController
}

func CreateServer() Server {
	LoadEnv()

	port := GetEnvValue("PORT")

	database, err := db.ConnectDB(
		GetEnvValue("DATABASE_USERNAME"),
		GetEnvValue("DATABASE_PASSWORD"),
		GetEnvValue("DATABASE_PORT"),
		GetEnvValue("DATABASE_NAME"),
	)

	if err != nil {
		panic(err.Error())
	}
	r := gin.Default()
	server := Server{
		R: r,
		FFController: controllers.FFController{
			DB: database,
		},
		EnvController: controllers.EnvController{
			DB: database,
		},
	}
	server.CreateFFSGroup()
	server.CreateEnvGroup()
	r.Run(port)

	return server
}

func ErrorHandler(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"message": message,
	})
}

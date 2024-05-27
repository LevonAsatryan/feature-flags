package server

import (
	"github.com/LevonAsatryan/feature-flags/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	DB *db.DB
	R  *gin.Engine
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
		DB: database,
		R:  r,
	}
	server.CreateFFSGroup()
	server.CreateEnvGroup()
	r.Run(port)

	return server
}

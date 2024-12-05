package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LevonAsatryan/feature-flags/tree/main/server/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Example struct {
	Id   primitive.ObjectID `json:"id"`
	Name string             `json:"name"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dbClient := db.ConnectDB()

	err := db.InsertMockData(dbClient)

	if err != nil {
		fmt.Errorf("Failed to insert into database with error:%s", err.Error())
		panic(err)
	}

	// Server part
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:3030")
}

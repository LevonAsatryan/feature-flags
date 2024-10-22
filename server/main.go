package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Example struct {
	Id   primitive.ObjectID `json:"id"`
	Name string             `json:"name"`
}

func main() {
	fmt.Println("hello motherfucker")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:3030")
}

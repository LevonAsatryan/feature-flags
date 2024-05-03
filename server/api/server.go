package server

import (
	"net/http"
	"strconv"

	"github.com/LevonAsatryan/feature-flags/api/controllers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port  int64
	store *Store
}

func NewServer(port int64) *Server {
	return &Server{port: port}
}

func (s *Server) Start() error {
	router := gin.Default()
	router.GET("/ping", handlePing)
	controllers.AddHandlers(router)
	err := router.Run("localhost:" + strconv.FormatInt(s.port, 10))
	return err
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

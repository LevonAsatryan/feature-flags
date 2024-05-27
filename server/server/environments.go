package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateEnvBody struct {
	Name string `json:"name"`
}

func (s *Server) CreateEnvGroup() *gin.RouterGroup {
	group := s.R.Group("/envs")
	group.GET("/", s.getEnvAll)
	group.GET("/:id", s.getEnvById)
	group.POST("/", s.createEnv)
	group.DELETE("/:id", s.deleteEnv)
	group.PUT("/:id", s.updateEnv)
	return group
}

func (s *Server) getEnvAll(c *gin.Context) {
	envs, err := s.DB.GetEnvAll()

	if err != nil {
		errorHandler(c, http.StatusInternalServerError, messages["intError"])
		return
	}

	c.JSON(http.StatusOK, envs)
}

func (s *Server) getEnvById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorHandler(c, http.StatusBadRequest, messages["invalidId"])
		return
	}

	env, err := s.DB.GetEnvById(id)

	if err != nil {
		errorHandler(c, http.StatusBadRequest, messages["envNotFound"])
		return
	}

	c.JSON(http.StatusOK, env)
}

func (s *Server) createEnv(c *gin.Context) {
	var rb CreateEnvBody

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&rb)

	if err != nil {
		errorHandler(c, http.StatusBadRequest, messages["badRequest"])
		return
	}

	fmt.Println(rb.Name)

	env, err := s.DB.CreateEnv(rb.Name)

	if err != nil {
		errorHandler(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, env)
}

func (s *Server) deleteEnv(c *gin.Context) {
	panic("NOT IMPLEMENTED")
}

func (s *Server) updateEnv(c *gin.Context) {
	panic("NOT IMPLEMENTED")
}

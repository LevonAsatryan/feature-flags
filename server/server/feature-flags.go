package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Env struct {
	ID   int    `json:"id"`
	Name string `json"name"`
}

type CreateFFBody struct {
	Name string `json:"name"`
	Env  string `json:"env"`
}

type UpdateFFBody struct {
	Name  string `json:"name"`
	Value bool   `json:"value"`
}

func (s *Server) CreateFFSGroup() *gin.RouterGroup {
	group := s.R.Group("/ffs")
	createMessages()
	group.GET("/", s.getFFAll)
	group.GET("/:id", s.getFFById)
	group.POST("/", s.createFF)
	group.DELETE("/:id", s.deleteFF)
	group.PUT("/:id", s.updateFF)
	return group
}

func (s *Server) getFFAll(c *gin.Context) {
	ffs, err := s.DB.GetFFAll()

	if err != nil {
		errorHandler(c, http.StatusInternalServerError, messages["intError"])
		return
	}

	c.JSON(http.StatusOK, ffs)
}

func (s *Server) createFF(c *gin.Context) {
	var rb CreateFFBody
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&rb)

	if err != nil {
		errorHandler(c, http.StatusBadRequest, err.Error())
	}

	ff, err := s.DB.CreateFF(rb.Name, rb.Env)

	if err != nil {
		errorHandler(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, &ff)
}

func (s *Server) updateFF(c *gin.Context) {
	var rb UpdateFFBody
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&rb)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorHandler(c, http.StatusBadRequest, messages["invalidId"])
		return
	}

	if err != nil {
		errorHandler(c, http.StatusBadRequest, messages["badRequest"])
		return
	}

	ff, err := s.DB.UpdateFF(id, rb.Name, rb.Value)

	if err != nil {
		errorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ff)
}

func (s *Server) getFFById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorHandler(c, http.StatusBadRequest, messages["invalidId"])
		return
	}
	ff, err := s.DB.GetFFById(id)

	if err != nil {
		errorHandler(c, http.StatusNotFound, messages["ffNotFound"])
		return
	}

	c.JSON(http.StatusOK, ff)
}

func (s *Server) deleteFF(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		errorHandler(c, http.StatusBadRequest, messages["invalidId"])
		return
	}

	_, err = s.DB.GetFFById(id)

	if err != nil {
		errorHandler(c, http.StatusBadRequest, messages["ffNotFound"])
		return
	}

	err = s.DB.DeleteFF(id)

	if err != nil {
		errorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func errorHandler(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"message": message,
	})
}

var messages map[string]string = make(map[string]string)

func createMessages() {
	messages["ffNotFound"] = "Feature flag with given ID was not found"
	messages["badRequest"] = "Bad request"
	messages["invalidId"] = "Provided ID is not valid"
	messages["intError"] = "Something went wrong"
	messages["envNotFound"] = "Environment with given ID was not found"
}

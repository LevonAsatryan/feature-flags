package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateFFSGroup() *gin.RouterGroup {
	group := s.R.Group("/ffs")
	group.GET("/", s.getFFAll)
	group.GET("/:id", s.getFFById)
	group.POST("/", s.createFF)
	group.DELETE("/:id", s.deleteFF)
	group.PUT("/:id", s.updateFF)
	group.PUT("/name", s.updateFFName)
	return group
}

func (s *Server) getFFAll(c *gin.Context) {
	ffs, err := s.FFController.GetAll(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, ffs)
}

func (s *Server) createFF(c *gin.Context) {
	ffs, err := s.FFController.CreateFF(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, ffs)
}

func (s *Server) getFFById(c *gin.Context) {
	ff, err := s.FFController.GetById(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, ff)
}

func (s *Server) updateFF(c *gin.Context) {
	ff, err := s.FFController.Update(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, ff)
}

func (s *Server) updateFFName(c *gin.Context) {
	ffs, err := s.FFController.UpdateName(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, ffs)
}

func (s *Server) deleteFF(c *gin.Context) {
	err := s.FFController.Delete(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

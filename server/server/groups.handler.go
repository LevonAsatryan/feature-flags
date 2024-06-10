package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateGroupsGroup() *gin.RouterGroup {
	group := s.R.Group("/group")
	group.GET("/", s.getGroupsAll)
	// group.GET("/:id", s.getFFById)
	group.POST("/", s.createGroup)
	group.PUT("/:id/addFF", s.addFF)
	// group.DELETE("/:id", s.deleteFF)
	// group.PUT("/:id", s.updateFF)
	// group.PUT("/name", s.updateFFName)
	return group
}

func (s *Server) getGroupsAll(c *gin.Context) {
	groups, err := s.GroupsController.GetAll(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, groups)
}

func (s *Server) createGroup(c *gin.Context) {
	envs, err := s.EnvController.GetAll(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	groups, err := s.GroupsController.Create(c, envs)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, groups)
}

func (s *Server) addFF(c *gin.Context) {
	group, err := s.GroupsController.GetById(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	fmt.Println(group.ID)

	err = s.FFController.AddFFToGroup(c, group.ID)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, group)
}

package server

import (
	"net/http"

	"github.com/LevonAsatryan/feature-flags/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

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
	envs, err := s.EnvController.GetAll(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, envs)
}

func (s *Server) getEnvById(c *gin.Context) {
	env, err := s.EnvController.GetById(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, env)
}

func (s *Server) createEnv(c *gin.Context) {
	envs, err := s.EnvController.GetAll(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	envToCopyFrom := envs[0]

	env, err := s.EnvController.Create(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	ffs, err := s.FFController.GetByEnvId(c, envToCopyFrom.ID)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	for _, ff := range ffs {
		_, createErr := s.FFController.DB.CreateFF(s.FFController.Ctx, db.CreateFFParams{
			Name: ff.Name,
			EnvID: pgtype.Int4{
				Int32: envToCopyFrom.ID,
				Valid: true,
			},
		})

		if createErr != nil {
			ErrorHandler(c, http.StatusInternalServerError, "could not create ffs for the environment")
			return
		}
	}

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, env)
}

func (s *Server) deleteEnv(c *gin.Context) {
	err := s.EnvController.Delete(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (s *Server) updateEnv(c *gin.Context) {
	err := s.EnvController.Update(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}

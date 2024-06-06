package server

import (
	"fmt"
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
		count, countErr := s.FFController.DB.CountFFByNameAndEnvId(
			s.FFController.Ctx,
			db.CountFFByNameAndEnvIdParams{
				Name: ff.Name,
				EnvID: pgtype.Int4{
					Int32: env.ID,
					Valid: true,
				},
			},
		)

		if countErr != nil {
			ErrorHandler(
				c,
				http.StatusInternalServerError,
				fmt.Sprintf(
					"could not get the count for ff by name %s and envID %d",
					ff.Name.String,
					env.ID,
				),
			)
			return
		}

		if count != 0 {
			continue
		}

		_, createErr := s.FFController.DB.CreateFF(s.FFController.Ctx, db.CreateFFParams{
			Name: ff.Name,
			EnvID: pgtype.Int4{
				Int32: env.ID,
				Valid: true,
			},
		})

		if createErr != nil {
			ErrorHandler(c, http.StatusInternalServerError, "could not create ffs for the environment")
			return
		}
	}

	groups, err := s.GroupsController.GetAll(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	for _, group := range groups {
		count, countErr := s.FFController.DB.CountGroupByNameAndEnvID(
			s.FFController.Ctx,
			db.CountGroupByNameAndEnvIDParams{
				Name: group.Name,
				EnvID: pgtype.Int4{
					Int32: env.ID,
					Valid: true,
				},
			},
		)

		if countErr != nil {
			ErrorHandler(
				c,
				http.StatusInternalServerError,
				fmt.Sprintf(
					"could not get the count for ff by name %s and envID %d",
					group.Name.String,
					env.ID,
				),
			)
			return
		}

		if count != 0 {
			continue
		}

		_, groupErr := s.GroupsController.DB.CreateGroup(c, db.CreateGroupParams{
			Name:  pgtype.Text{String: group.Name.String, Valid: true},
			EnvID: pgtype.Int4{Int32: env.ID, Valid: true},
		})

		if groupErr != nil {
			ErrorHandler(c, http.StatusInternalServerError, "could not create group for the environment")
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
	count, err := s.EnvController.GetEnvCount(c)

	if err != nil {
		ErrorHandler(c, err.Code, err.Err.Error())
		return
	}

	if count == 1 {
		ErrorHandler(c, http.StatusBadRequest, "can not delete the last environment")
		return
	}

	err = s.EnvController.Delete(c)

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

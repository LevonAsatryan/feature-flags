package server

import (
	"context"
	"fmt"

	"github.com/LevonAsatryan/feature-flags/controllers"
	"github.com/LevonAsatryan/feature-flags/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	DB            *db.Queries
	DBTX          db.DBTX
	R             *gin.Engine
	FFController  controllers.FFController
	EnvController controllers.EnvController
}

func CreateServer() Server {
	LoadEnv()

	port := GetEnvValue("PORT")
	user := GetEnvValue("DATABASE_USERNAME")
	dbName := GetEnvValue("DATABASE_NAME")
	dbPassword := GetEnvValue("DATABASE_PASSWORD")

	ctx := context.Background()
	conn, err := pgx.Connect(
		ctx,
		fmt.Sprintf(
			"user=%s dbname=%s password=%s host=localhost sslmode=disable",
			user,
			dbName,
			dbPassword,
		),
	)

	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	database := db.New(conn)

	r := gin.Default()
	server := Server{
		R: r,
		FFController: controllers.FFController{
			DB:  database,
			Ctx: ctx,
		},
		EnvController: controllers.EnvController{
			DB:  database,
			Ctx: ctx,
		},
	}
	err = server.CreateTables(conn)
	if err != nil {
		panic(err)
	}
	server.EnvController.CheckAndCreateEnv()
	server.CreateFFSGroup()
	server.CreateEnvGroup()
	r.Run(port)

	return server
}

func (s *Server) CreateTables(conn db.DBTX) error {
	err := s.createEnvTable(conn)
	if err != nil {
		return err
	}
	err = s.createFFTable(conn)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) createEnvTable(conn db.DBTX) error {
	_, err := conn.Exec(s.EnvController.Ctx, `
		CREATE TABLE if NOT EXISTS env (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) UNIQUE,
			origin_host VARCHAR(255) UNIQUE,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`)
	return err
}

func (s *Server) createFFTable(conn db.DBTX) error {
	_, err := conn.Exec(s.FFController.Ctx, `
		CREATE TABLE if NOT EXISTS feature_flags (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			value BOOLEAN DEFAULT true,
			env_id integer REFERENCES env (id),
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)

	return err
}

func ErrorHandler(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"message": message,
	})
}

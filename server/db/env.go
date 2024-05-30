package db

import (
	"fmt"
	"net/http"

	"github.com/LevonAsatryan/feature-flags/types"
)

type Env struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Created_at string `db:"created_at" json:"createdAt"`
	Updated_at string `db:"updated_at" json:"updatedAt"`
}

func (db *DB) GetEnvById(id int) (*Env, *types.Error) {
	var env Env

	res := db.Conn.QueryRow(
		"SELECT * FROM env WHERE id = $1", id,
	)

	err := res.Scan(&env.ID, &env.Name, &env.Created_at, &env.Updated_at)

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("no environment with id=%d exists", id),
		}
	}

	return &env, nil
}

func (db *DB) CreateEnv(name string) (*Env, *types.Error) {
	var envID int
	res := db.Conn.QueryRow(
		"INSERT INTO env (name) VALUES ($1) RETURNING id", name,
	)

	err := res.Scan(&envID)

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to create env with name %s", name),
		}
	}

	env, tErr := db.GetEnvById(envID)

	if tErr != nil {
		return nil, tErr
	}

	return env, nil
}

func (db *DB) GetEnvAll() ([]Env, *types.Error) {
	var envs []Env

	res, err := db.Conn.Query(
		"SELECT * FROM env",
	)

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to fetch the environments"),
		}
	}

	for res.Next() {
		var env Env
		err = res.Scan(&env.ID, &env.Name, &env.Created_at, &env.Updated_at)
		if err != nil {
			return nil, &types.Error{
				Code: http.StatusInternalServerError,
				Err:  fmt.Errorf("failed to attach environment from db to struct"),
			}
		}
		envs = append(envs, env)
	}

	return envs, nil
}

func (db *DB) UpdateEnv(id int, name string) (*Env, *types.Error) {
	_, err := db.Conn.Query(
		"UPDATE env SET name = $1 WHERE id = $2", name, id,
	)

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to update environment by id=%d and name=%s", id, name),
		}
	}

	env, tErr := db.GetEnvById(id)

	return env, tErr
}

func (db *DB) DeleteEnv(id int) *types.Error {
	query := "DELETE FROM env WHERE id=$1"

	_, err := db.Conn.Query(query, id)

	if err != nil {
		return &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("failed to delete environment by id=%d", id),
		}
	}

	return nil
}

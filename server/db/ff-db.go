package db

import (
	"fmt"
	"net/http"

	"github.com/LevonAsatryan/feature-flags/types"
)

type FeatureFlag struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Value      bool   `db:"value" json:"value"`
	Env        int    `db:"env_id" json:"env"`
	Created_at string `db:"created_at" json:"createdAt"`
	Updated_at string `db:"updated_at" json:"updatedAt"`
}

func (db *DB) CreateFF(name string) ([]int, *types.Error) {

	envs, tErr := db.GetEnvAll()

	if tErr != nil {
		return nil, tErr
	}

	var ffIDs []int

	for _, env := range envs {
		var ffID int
		res := db.Conn.QueryRow(
			"INSERT INTO feature_flags (name, env_id) VALUES ($1, $2) RETURNING id",
			name,
			env.ID,
		)

		err := res.Scan(&ffID)

		if err != nil {
			return nil, &types.Error{
				Code: http.StatusInternalServerError,
				Err:  fmt.Errorf("failed to insert the feature flag"),
			}
		}

		ffIDs = append(ffIDs, ffID)
	}

	return ffIDs, nil
}

func (db *DB) UpdateFF(id int, value bool) (*FeatureFlag, *types.Error) {
	query := "UPDATE feature_flags SET value = $1 WHERE id = $2"
	_, err := db.Conn.Query(query, value, id)

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("failed to update feature flag with id=%s", id),
		}
	}

	ff, tErr := db.GetFFById(id)

	if tErr != nil {
		return nil, tErr
	}

	return ff, nil
}

func (db *DB) GetFFByEnvId(envId int) ([]FeatureFlag, *types.Error) {
	query := "SELECT * FROM feature_flags WHERE env_id=$1"

	var featureFlags []FeatureFlag

	res, err := db.Conn.Query(query, envId)

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to fetch feature flags from db"),
		}
	}

	for res.Next() {
		var ff FeatureFlag

		err = res.Scan(
			&ff.ID,
			&ff.Name,
			&ff.Value,
			&ff.Env,
			&ff.Created_at,
			&ff.Updated_at,
		)

		if err != nil {
			return nil, &types.Error{
				Code: http.StatusInternalServerError,
				Err:  err,
			}
		}

		featureFlags = append(featureFlags, ff)
	}

	return featureFlags, nil
}

func (db *DB) CopyFromExisting(existingEnvId, destEnvId int) *types.Error {
	ffs, tErr := db.GetFFByEnvId(existingEnvId)

	if tErr != nil {
		return tErr
	}

	query := `INSERT INTO feature_flags 
		(name, env_id, value) VALUES `

	separator := ","

	for index, ff := range ffs {
		if index == len(ffs)-1 {
			separator = ";"
		}
		query += fmt.Sprintf("('%s', %d, %v)%s ", ff.Name, destEnvId, ff.Value, separator)
	}

	_, err := db.Conn.Query(query)

	if err != nil {
		return &types.Error{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	return nil
}

func (db *DB) GetFFById(id int) (*FeatureFlag, *types.Error) {
	row := db.Conn.QueryRow(
		`SELECT * FROM feature_flags WHERE id = $1`, id,
	)

	var featureFlag FeatureFlag

	err := row.Scan(
		&featureFlag.ID,
		&featureFlag.Name,
		&featureFlag.Value,
		&featureFlag.Env,
		&featureFlag.Created_at,
		&featureFlag.Updated_at,
	)

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("feature flag with id=%d, not found", id),
		}
	}

	return &featureFlag, nil
}

func (db *DB) GetFFAll() ([]FeatureFlag, *types.Error) {
	var ffs []FeatureFlag
	rows, err := db.Conn.Query("SELECT * FROM feature_flags")

	if err != nil {
		return nil, &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to fetch feature flags"),
		}
	}

	for rows.Next() {
		var ff FeatureFlag
		rows.Scan(
			&ff.ID,
			&ff.Name,
			&ff.Value,
			&ff.Env,
			&ff.Created_at,
			&ff.Updated_at,
		)
		ffs = append(ffs, ff)
	}

	if ffs == nil {
		ffs = make([]FeatureFlag, 0)
	}

	return ffs, nil
}

func (db *DB) DeleteFF(id int) *types.Error {
	_, err := db.Conn.Query(`DELETE FROM feature_flags WHERE id = $1`, id)

	if err != nil {
		return &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to delete a feature flag with id=%d", id),
		}
	}

	return nil
}

func (db *DB) DeleteFFsByEnvId(id int) *types.Error {
	_, err := db.Conn.Query(`DELETE FROM feature_flags WHERE env_id = $1`, id)

	if err != nil {
		return &types.Error{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("failed to delete a feature flag with env_id=%d", id),
		}
	}

	return nil
}

package db

import (
	"fmt"
	"strconv"
)

type FeatureFlag struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Value      bool   `json:"value"`
	Env        int    `json:"env"`
	Created_at string `json:"createdAt"`
	Updated_at string `json:"updatedAt"`
}

func (db *DB) createFF(name, envName string) (*FeatureFlag, error) {
	env, err := db.GetEnvByName(envName)

	if err != nil {
		return nil, err
	}

	envID := strconv.Itoa(env.ID)
	rows, err := db.Conn.Query(
		fmt.Sprintf(
			"INSERT INTO feature-flags (name, value, env) VALUES (%s, %s, %s)",
			name,
			strconv.FormatBool(true),
			envID,
		),
	)

	if err != nil {
		return nil, err
	}
	var ff FeatureFlag
	rows.Scan(&ff)

	return &ff, nil
}

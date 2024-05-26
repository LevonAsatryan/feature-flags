package db

import "fmt"

func (db *DB) GetEnvByName(name string) (*Env, error) {
	var env Env

	res, err := db.Conn.Query(
		fmt.Sprintf("SELECT * FROM env WHERE name=%s", name),
	)

	if err != nil {
		return nil, err
	}

	res.Scan(&env)

	return &env, nil
}

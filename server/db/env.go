package db

func (db *DB) GetEnvByName(name string) (*Env, error) {
	var env Env

	res := db.Conn.QueryRow(
		"SELECT * FROM env WHERE name = $1", name,
	)

	err := res.Scan(&env.ID, &env.Name, &env.Created_at, &env.Updated_at)

	if err != nil {
		return nil, err
	}

	// if env.Name == "" {
	// 	return nil, fmt.Errorf("the environment with name %s not found", name)
	// }

	return &env, nil
}

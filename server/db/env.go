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

	return &env, nil
}

func (db *DB) GetEnvById(id int) (*Env, error) {
	var env Env

	res := db.Conn.QueryRow(
		"SELECT * FROM env WHERE id = $1", id,
	)

	err := res.Scan(&env.ID, &env.Name, &env.Created_at, &env.Updated_at)

	if err != nil {
		return nil, err
	}

	return &env, nil
}

func (db *DB) CreateEnv(name string) (*Env, error) {
	var envID int
	res := db.Conn.QueryRow(
		"INSERT INTO env (name) VALUES ($1) RETURNING id", name,
	)

	err := res.Scan(&envID)

	if err != nil {
		return nil, err
	}

	env, err := db.GetEnvById(envID)

	if err != nil {
		return nil, err
	}

	return env, err
}

func (db *DB) GetEnvAll() ([]Env, error) {
	var envs []Env

	res, err := db.Conn.Query(
		"SELECT * FROM env",
	)

	if err != nil {
		return nil, err
	}

	for res.Next() {
		var env Env
		res.Scan(&env.ID, &env.Name, &env.Created_at, &env.Updated_at)
		envs = append(envs, env)
	}

	return envs, nil
}

func (db *DB) UpdateEnv(id int, name string) (*Env, error) {
	_, err := db.GetEnvById(id)

	if err != nil {
		return nil, err
	}

	res := db.Conn.QueryRow(
		"UPDATE env SET name = $1 WHERE id = $2", name, id,
	)

	var env Env

	err = res.Scan(&env.ID, &env.Name, &env.Created_at, &env.Updated_at)

	if err != nil {
		return nil, err
	}

	return &env, nil
}

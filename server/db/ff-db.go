package db

type FeatureFlag struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Value      bool   `db:"value" json:"value"`
	Env        int    `db:"env_id" json:"env"`
	Created_at string `db:"created_at" json:"createdAt"`
	Updated_at string `db:"updated_at" json:"updatedAt"`
}

func (db *DB) CreateFF(name, envName string) (*FeatureFlag, error) {
	env, err := db.GetEnvByName(envName)

	if err != nil {
		return nil, err
	}
	var ffID int

	rows := db.Conn.QueryRow(
		"INSERT INTO feature_flags (name, env_id) VALUES ($1, $2) RETURNING id",
		name,
		env.ID,
	)

	if err != nil {
		return nil, err
	}
	// var ff FeatureFlag
	err = rows.Scan(&ffID)

	if err != nil {
		return nil, err
	}

	ff, err := db.GetFFById(ffID)

	if err != nil {
		return nil, err
	}

	return ff, nil
}

func (db *DB) UpdateFF(id int, name string, value bool) (*FeatureFlag, error) {
	query := "UPDATE feature_flags SET name = $1, value = $2 WHERE id = $3"
	_, err := db.Conn.Query(query, name, value, id)

	if err != nil {
		return nil, err
	}

	ff, err := db.GetFFById(id)

	if err != nil {
		return nil, err
	}

	return ff, nil
}

func (db *DB) GetFFById(id int) (*FeatureFlag, error) {
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
		return nil, err
	}

	return &featureFlag, nil
}

func (db *DB) GetFFAll() ([]FeatureFlag, error) {
	var ffs []FeatureFlag
	rows, err := db.Conn.Query("SELECT * FROM feature_flags")

	if err != nil {
		return nil, err
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

	return ffs, nil
}

func (db *DB) DeleteFF(id int) error {
	_, err := db.Conn.Query(`DELETE FROM feature_flags WHERE id = $1`, id)

	if err != nil {
		return err
	}

	return nil
}

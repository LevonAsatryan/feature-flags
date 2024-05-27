package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	ConnStr string
	Conn    *sql.DB
}

type Env struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

var db DB

func ConnectDB(username, password, port, name string) (*DB, error) {
	db = DB{
		ConnStr: fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", username, password, port, name),
	}

	fmt.Println(db.ConnStr)

	conn, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		return nil, err
	}

	db.Conn = conn
	err = db.createTables()
	if err != nil {
		return nil, err
	}
	return &db, nil
}

func (db *DB) createTables() error {
	err := db.createEnvTable()
	if err != nil {
		return err
	}
	err = db.createFFTable()

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) createEnvTable() error {
	_, err := db.Conn.Query(
		"CREATE TABLE if NOT EXISTS env (id SERIAL PRIMARY KEY, name VARCHAR(255) UNIQUE, created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW())",
	)

	if err != nil {
		return err
	}
	var rowCount int

	res := db.Conn.QueryRow(
		"SELECT COUNT(*) FROM env",
	)

	if err != nil {
		return err
	}

	res.Scan(&rowCount)

	if rowCount == 0 {
		_, err = db.Conn.Query(
			"INSERT INTO env (name) VALUES ('dev')",
		)
	}
	return err
}

func (db *DB) createFFTable() error {
	_, err := db.Conn.Query(
		"CREATE TABLE if NOT EXISTS feature_flags (id SERIAL PRIMARY KEY, name VARCHAR(255), value BOOLEAN, env_id integer REFERENCES env (id), created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW())",
	)

	return err
}

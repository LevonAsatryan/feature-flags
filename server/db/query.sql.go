// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createEnv = `-- name: CreateEnv :one
INSERT INTO env (name, origin_host) VALUES ($1, $2) RETURNING id, name, origin_host, created_at, updated_at
`

type CreateEnvParams struct {
	Name       pgtype.Text
	OriginHost pgtype.Text
}

func (q *Queries) CreateEnv(ctx context.Context, arg CreateEnvParams) (Env, error) {
	row := q.db.QueryRow(ctx, createEnv, arg.Name, arg.OriginHost)
	var i Env
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OriginHost,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createFF = `-- name: CreateFF :one
INSERT INTO feature_flags (name, env_id) VALUES ($1, $2) RETURNING id, name, value, env_id, created_at, updated_at
`

type CreateFFParams struct {
	Name  pgtype.Text
	EnvID pgtype.Int4
}

func (q *Queries) CreateFF(ctx context.Context, arg CreateFFParams) (FeatureFlag, error) {
	row := q.db.QueryRow(ctx, createFF, arg.Name, arg.EnvID)
	var i FeatureFlag
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Value,
		&i.EnvID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteEnv = `-- name: DeleteEnv :exec
DELETE FROM env WHERE id = $1
`

func (q *Queries) DeleteEnv(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteEnv, id)
	return err
}

const deleteFF = `-- name: DeleteFF :exec
DELETE FROM feature_flags WHERE id = $1
`

func (q *Queries) DeleteFF(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteFF, id)
	return err
}

const deleteFFsByEnvId = `-- name: DeleteFFsByEnvId :exec
DELETE FROM feature_flags WHERE env_id = $1
`

func (q *Queries) DeleteFFsByEnvId(ctx context.Context, envID pgtype.Int4) error {
	_, err := q.db.Exec(ctx, deleteFFsByEnvId, envID)
	return err
}

const getEnv = `-- name: GetEnv :one
SELECT id, name, origin_host, created_at, updated_at FROM env WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEnv(ctx context.Context, id int32) (Env, error) {
	row := q.db.QueryRow(ctx, getEnv, id)
	var i Env
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OriginHost,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getEnvAll = `-- name: GetEnvAll :many
SELECT id, name, origin_host, created_at, updated_at FROM env
`

func (q *Queries) GetEnvAll(ctx context.Context) ([]Env, error) {
	rows, err := q.db.Query(ctx, getEnvAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Env
	for rows.Next() {
		var i Env
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.OriginHost,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFeatureFlag = `-- name: GetFeatureFlag :one
SELECT id, name, value, env_id, created_at, updated_at FROM feature_flags WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFeatureFlag(ctx context.Context, id int32) (FeatureFlag, error) {
	row := q.db.QueryRow(ctx, getFeatureFlag, id)
	var i FeatureFlag
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Value,
		&i.EnvID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFeatureFlagAll = `-- name: GetFeatureFlagAll :many
SELECT id, name, value, env_id, created_at, updated_at FROM feature_flags
`

func (q *Queries) GetFeatureFlagAll(ctx context.Context) ([]FeatureFlag, error) {
	rows, err := q.db.Query(ctx, getFeatureFlagAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeatureFlag
	for rows.Next() {
		var i FeatureFlag
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Value,
			&i.EnvID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFeatureFlagByEnvId = `-- name: GetFeatureFlagByEnvId :many
SELECT id, name, value, env_id, created_at, updated_at FROM feature_flags WHERE env_id = $1
`

func (q *Queries) GetFeatureFlagByEnvId(ctx context.Context, envID pgtype.Int4) ([]FeatureFlag, error) {
	rows, err := q.db.Query(ctx, getFeatureFlagByEnvId, envID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeatureFlag
	for rows.Next() {
		var i FeatureFlag
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Value,
			&i.EnvID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEnv = `-- name: UpdateEnv :exec
UPDATE env set name = $2, origin_host = $3 WHERE id = $1 RETURNING id, name, origin_host, created_at, updated_at
`

type UpdateEnvParams struct {
	ID         int32
	Name       pgtype.Text
	OriginHost pgtype.Text
}

func (q *Queries) UpdateEnv(ctx context.Context, arg UpdateEnvParams) error {
	_, err := q.db.Exec(ctx, updateEnv, arg.ID, arg.Name, arg.OriginHost)
	return err
}

const updateFF = `-- name: UpdateFF :exec
UPDATE feature_flags set value = $2 WHERE id = $1 RETURNING id, name, value, env_id, created_at, updated_at
`

type UpdateFFParams struct {
	ID    int32
	Value pgtype.Bool
}

func (q *Queries) UpdateFF(ctx context.Context, arg UpdateFFParams) error {
	_, err := q.db.Exec(ctx, updateFF, arg.ID, arg.Value)
	return err
}

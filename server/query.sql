-- name: GetEnv :one
SELECT * FROM env WHERE id = $1 LIMIT 1;

-- name: GetEnvCount :one
SELECT COUNT(*) FROM env;

-- name: GetEnvAll :many
SELECT * FROM env;

-- name: CreateEnv :one
INSERT INTO env (name, origin_host) VALUES ($1, $2) RETURNING *;

-- name: UpdateEnv :exec
UPDATE env set name = $2, origin_host = $3 WHERE id = $1 RETURNING *;

-- name: DeleteEnv :exec
DELETE FROM env WHERE id = $1;

-- name: GetFeatureFlag :one
SELECT * FROM feature_flags WHERE id = $1 LIMIT 1;

-- name: GetFeatureFlagAll :many
SELECT * FROM feature_flags;

-- name: GetFeatureFlagByEnvId :many
SELECT * FROM feature_flags WHERE env_id = $1;

-- name: CreateFF :one
INSERT INTO feature_flags (name, env_id) VALUES ($1, $2) RETURNING *;

-- name: UpdateFF :exec
UPDATE feature_flags set value = $2 WHERE id = $1 RETURNING *;

-- name: DeleteFF :exec
DELETE FROM feature_flags WHERE id = $1;

-- name: DeleteFFsByEnvId :exec
DELETE FROM feature_flags WHERE env_id = $1;
-- name: GetEnv :one
SELECT * FROM env WHERE id = $1 LIMIT 1;

-- name: GetEnvCount :one
SELECT COUNT(*) FROM env;

-- name: GetEnvAll :many
SELECT * FROM env;

-- name: CreateEnv :one
INSERT INTO env (name, origin_host) VALUES ($1, $2) RETURNING *;

-- name: UpdateEnv :exec
UPDATE env set name = $2, origin_host = $3 WHERE id = $1;

-- name: DeleteEnv :exec
DELETE FROM env WHERE id = $1;

-- name: GetFF :one
SELECT * FROM feature_flags WHERE id = $1 LIMIT 1;

-- name: GetFFAll :many
SELECT * FROM feature_flags;

-- name: GetFFByEnvId :many
SELECT * FROM feature_flags WHERE env_id = $1;

-- name: GetFFByName :many
SELECT * FROM feature_flags WHERE name = $1;

-- name: CountFFByNameAndEnvId :one
SELECT COUNT(*) FROM feature_flags WHERE name = $1 AND env_id = $2;

-- name: CreateFF :one
INSERT INTO feature_flags (name, env_id) VALUES ($1, $2) RETURNING *;

-- name: UpdateFF :exec
UPDATE feature_flags set value = $2 WHERE id = $1;

-- name: UpdateFFName :exec
UPDATE feature_flags set name = $2 WHERE name = $1;

-- name: DeleteFF :exec
DELETE FROM feature_flags WHERE id = $1;

-- name: DeleteFFsByEnvId :exec
DELETE FROM feature_flags WHERE env_id = $1;

-- name: CreateGroup :one
INSERT INTO groups (name, env_id) VALUES ($1, $2) RETURNING *;

-- name: GetGroupsAll :many
SELECT * FROM groups;

-- name: GetOne :one
SELECT * FROM groups WHERE id = $1;

-- name: CountGroupByNameAndEnvID :one
SELECT COUNT(*) FROM groups WHERE name = $1 AND env_id = $2;

-- name: UpdateGroup :exec
UPDATE groups set name=$1 WHERE name=$2;

-- name: DeleteGroup :exec
DELETE FROM groups WHERE id=$1;
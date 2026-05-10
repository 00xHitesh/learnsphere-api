-- name: ListFeatureFlags :many
SELECT * FROM feature_flags ORDER BY name;

-- name: GetFeatureFlagByName :one
SELECT * FROM feature_flags WHERE name = $1 LIMIT 1;

-- name: ToggleFeatureFlag :one
UPDATE feature_flags SET enabled = NOT enabled WHERE name = $1 RETURNING *;
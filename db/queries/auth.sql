-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = lower($1) LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (name, email, password_hash, role)
VALUES ($1, lower($2), $3, 'user')
RETURNING *;

-- name: UpdateUserPoints :one
UPDATE users SET points = points + $2 WHERE id = $1 RETURNING *;

-- name: UpdateUserLocale :one
UPDATE users SET region_code = $2, currency_code = $3 WHERE id = $1 RETURNING *;

-- name: GetUserInterests :many
SELECT category FROM user_interests WHERE user_id = $1 ORDER BY category;

-- name: DeleteUserInterests :exec
DELETE FROM user_interests WHERE user_id = $1;

-- name: InsertUserInterest :exec
INSERT INTO user_interests (user_id, category) VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: CreateResetToken :one
INSERT INTO password_reset_tokens (user_id, token, expires_at)
VALUES ($1, $2, NOW() + INTERVAL '1 hour')
RETURNING *;

-- name: GetActiveResetToken :one
SELECT * FROM password_reset_tokens
WHERE token = $1 AND used_at IS NULL AND expires_at > NOW()
LIMIT 1;

-- name: MarkResetTokenUsed :exec
UPDATE password_reset_tokens SET used_at = NOW() WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users SET password_hash = $2 WHERE id = $1;
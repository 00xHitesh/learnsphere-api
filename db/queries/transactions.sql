-- name: ListTransactionsByUser :many
SELECT * FROM transactions WHERE user_id = $1 ORDER BY created_at DESC;

-- name: ListAllTransactions :many
SELECT * FROM transactions ORDER BY created_at DESC;

-- name: InsertTransaction :one
INSERT INTO transactions (user_id, course_id, amount, status)
VALUES ($1, $2, $3, $4)
RETURNING *;
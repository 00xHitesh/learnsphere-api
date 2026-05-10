-- name: ListPublishedLiveSessions :many
SELECT * FROM live_sessions WHERE status = 'published' ORDER BY scheduled_at;

-- name: ListAllLiveSessions :many
SELECT * FROM live_sessions ORDER BY scheduled_at;

-- name: GetLiveSessionByID :one
SELECT * FROM live_sessions WHERE id = $1 LIMIT 1;

-- name: InsertLiveSession :one
INSERT INTO live_sessions (title, instructor, description, scheduled_at, duration_minutes, topic, join_url, course_id, registered_count, max_attendees, list_price, sale_price, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING *;

-- name: UpdateLiveSession :one
UPDATE live_sessions
SET title = $2, instructor = $3, description = $4, scheduled_at = $5,
    duration_minutes = $6, topic = $7, join_url = $8, course_id = $9,
    registered_count = $10, max_attendees = $11, list_price = $12,
    sale_price = $13, status = $14
WHERE id = $1
RETURNING *;

-- name: UpdateLiveSessionStatus :one
UPDATE live_sessions SET status = $2 WHERE id = $1 RETURNING *;

-- name: DeleteLiveSession :exec
DELETE FROM live_sessions WHERE id = $1;
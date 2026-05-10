-- name: GetOrCreateChatSession :one
INSERT INTO chat_sessions (user_id, course_id)
VALUES ($1, $2)
ON CONFLICT (user_id, course_id) DO UPDATE SET updated_at = NOW()
RETURNING *;

-- name: ListChatMessages :many
SELECT * FROM chat_messages WHERE session_id = $1 ORDER BY created_at ASC;

-- name: InsertChatMessage :one
INSERT INTO chat_messages (session_id, role, content)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteChatMessages :exec
DELETE FROM chat_messages WHERE session_id = $1;
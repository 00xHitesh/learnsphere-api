-- name: ListWishlistByUser :many
SELECT * FROM wishlists WHERE user_id = $1 ORDER BY added_at DESC;

-- name: GetWishlistItem :one
SELECT * FROM wishlists WHERE user_id = $1 AND course_id = $2 LIMIT 1;

-- name: InsertWishlistItem :exec
INSERT INTO wishlists (user_id, course_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: DeleteWishlistItem :exec
DELETE FROM wishlists WHERE user_id = $1 AND course_id = $2;
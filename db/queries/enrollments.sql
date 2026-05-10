-- name: ListEnrollmentsByUser :many
SELECT * FROM enrollments WHERE user_id = $1 ORDER BY enrolled_at DESC;

-- name: GetEnrollment :one
SELECT * FROM enrollments WHERE user_id = $1 AND course_id = $2 LIMIT 1;

-- name: InsertEnrollment :one
INSERT INTO enrollments (user_id, course_id)
VALUES ($1, $2)
ON CONFLICT (user_id, course_id) DO UPDATE SET updated_at = NOW()
RETURNING *;

-- name: GetCompletedLessonIDs :many
SELECT lesson_id FROM enrollment_lesson_progress WHERE enrollment_id = $1;

-- name: GetCompletedLessonIDsForEnrollments :many
SELECT enrollment_id, lesson_id FROM enrollment_lesson_progress
WHERE enrollment_id = ANY($1::uuid[]);

-- name: InsertLessonProgress :exec
INSERT INTO enrollment_lesson_progress (enrollment_id, lesson_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: DeleteLessonProgress :exec
DELETE FROM enrollment_lesson_progress WHERE enrollment_id = $1 AND lesson_id = $2;

-- name: AddEnrollmentTimeSpent :one
UPDATE enrollments SET time_spent_seconds = time_spent_seconds + $2
WHERE id = $1 RETURNING *;

-- name: UpdatePassedQuizIds :one
UPDATE enrollments SET passed_quiz_ids = $2 WHERE id = $1 RETURNING *;

-- name: IncrementUserPoints :exec
UPDATE users SET points = points + $2 WHERE id = $1;
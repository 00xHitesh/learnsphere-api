-- name: ListPublishedCourses :many
SELECT * FROM courses WHERE status = 'published' ORDER BY created_at DESC;

-- name: ListAllCourses :many
SELECT * FROM courses ORDER BY created_at DESC;

-- name: GetCourseBySlug :one
SELECT * FROM courses WHERE slug = $1 AND status = 'published' LIMIT 1;

-- name: GetCourseByID :one
SELECT * FROM courses WHERE id = $1 LIMIT 1;

-- name: InsertCourse :one
INSERT INTO courses (title, slug, description, summary, category, level, list_price, sale_price, instructor, instructor_bio, cover, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: UpdateCourse :one
UPDATE courses
SET title = $2, slug = $3, description = $4, summary = $5, category = $6,
    level = $7, list_price = $8, sale_price = $9, instructor = $10,
    instructor_bio = $11, cover = $12, status = $13
WHERE id = $1
RETURNING *;

-- name: UpdateCourseStatus :one
UPDATE courses SET status = $2 WHERE id = $1 RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM courses WHERE id = $1;

-- name: GetTagsByNames :many
SELECT * FROM tags WHERE name = ANY($1::text[]);

-- name: UpsertTag :one
INSERT INTO tags (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name RETURNING *;

-- name: DeleteCourseTagsByCourse :exec
DELETE FROM course_tags WHERE course_id = $1;

-- name: InsertCourseTag :exec
INSERT INTO course_tags (course_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: GetTagNamesByCourseID :many
SELECT t.name FROM course_tags ct JOIN tags t ON t.id = ct.tag_id WHERE ct.course_id = $1 ORDER BY t.name;

-- name: GetTagNamesForCourseIDs :many
SELECT ct.course_id, t.name
FROM course_tags ct JOIN tags t ON t.id = ct.tag_id
WHERE ct.course_id = ANY($1::uuid[])
ORDER BY ct.course_id, t.name;

-- name: ListLessonsByCourse :many
SELECT * FROM lessons WHERE course_id = $1 ORDER BY sort_order;

-- name: ListReviewsByCourse :many
SELECT * FROM reviews WHERE course_id = $1 ORDER BY reviewed_at DESC;

-- name: GetQuizByCourse :one
SELECT * FROM quizzes WHERE course_id = $1 LIMIT 1;

-- name: ListQuizQuestions :many
SELECT * FROM quiz_questions WHERE quiz_id = $1 ORDER BY sort_order;

-- name: ListQuizOptions :many
SELECT * FROM quiz_options WHERE question_id = ANY($1::uuid[]) ORDER BY question_id, sort_order;
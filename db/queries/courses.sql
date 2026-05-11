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

-- name: InsertLesson :one
INSERT INTO lessons (course_id, title, duration, video_url, is_preview, description, sort_order)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateLesson :one
UPDATE lessons
SET title = $3, duration = $4, video_url = $5, is_preview = $6, description = $7, sort_order = $8
WHERE id = $1 AND course_id = $2
RETURNING *;

-- name: DeleteLesson :exec
DELETE FROM lessons WHERE id = $1 AND course_id = $2;

-- name: InsertQuiz :one
INSERT INTO quizzes (course_id, title, description, passing_score)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateQuiz :one
UPDATE quizzes
SET title = $2, description = $3, passing_score = $4
WHERE id = $1
RETURNING *;

-- name: DeleteQuiz :exec
DELETE FROM quizzes WHERE id = $1;

-- name: InsertQuizQuestion :one
INSERT INTO quiz_questions (quiz_id, prompt, correct_option_keys, is_multi_select, explanation, sort_order)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: DeleteQuizQuestions :exec
DELETE FROM quiz_questions WHERE quiz_id = $1;

-- name: InsertQuizOption :exec
INSERT INTO quiz_options (question_id, option_key, text, sort_order)
VALUES ($1, $2, $3, $4);
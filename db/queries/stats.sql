-- name: GetAdminOverview :one
SELECT
  (SELECT COUNT(*) FROM users WHERE role = 'user')::BIGINT             AS total_users,
  (SELECT COUNT(*) FROM courses WHERE status = 'published')::BIGINT    AS total_courses,
  (SELECT COUNT(*) FROM enrollments)::BIGINT                           AS total_enrollments,
  COALESCE((SELECT SUM(amount) FROM transactions WHERE status = 'success'), 0)::NUMERIC AS revenue;

-- name: GetTopCourses :many
SELECT * FROM courses WHERE status = 'published' ORDER BY student_count DESC LIMIT 5;

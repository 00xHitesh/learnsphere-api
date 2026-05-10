-- LearnSphere AI — PostgreSQL 16+ Schema

CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "unaccent";

-- ENUMs
CREATE TYPE user_role          AS ENUM ('user', 'admin');
CREATE TYPE course_level       AS ENUM ('Beginner', 'Intermediate', 'Advanced');
CREATE TYPE transaction_status AS ENUM ('success', 'failed');
CREATE TYPE chat_role          AS ENUM ('user', 'assistant');
CREATE TYPE content_status     AS ENUM ('published', 'draft', 'in_review', 'archived');
CREATE TYPE notification_type  AS ENUM ('enrollment','completion','certificate','quiz_passed','quiz_failed','new_course','announcement','review_posted');
CREATE TYPE discount_type      AS ENUM ('percent', 'fixed');
CREATE TYPE session_status     AS ENUM ('upcoming', 'live', 'ended');

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER LANGUAGE plpgsql AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$;

-- currencies before regions (FK dependency)
CREATE TABLE currencies (
  code         TEXT          PRIMARY KEY,
  symbol       TEXT          NOT NULL,
  name         TEXT          NOT NULL,
  rate_vs_usd  NUMERIC(14,6) NOT NULL DEFAULT 1 CHECK (rate_vs_usd > 0)
);

CREATE TABLE regions (
  code                  TEXT PRIMARY KEY,
  name                  TEXT NOT NULL,
  default_currency_code TEXT NOT NULL REFERENCES currencies (code)
);

CREATE TABLE users (
  id            UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  name          TEXT        NOT NULL,
  email         TEXT        NOT NULL UNIQUE,
  password_hash TEXT        NOT NULL,
  role          user_role   NOT NULL DEFAULT 'user',
  avatar_url    TEXT,
  bio           TEXT,
  points        INTEGER     NOT NULL DEFAULT 0,
  region_code   TEXT        NOT NULL DEFAULT 'US',
  currency_code TEXT        NOT NULL DEFAULT 'USD',
  created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_role  ON users (role);

CREATE TRIGGER trg_users_updated_at
  BEFORE UPDATE ON users
  FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TABLE user_interests (
  user_id  UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  category TEXT NOT NULL,
  PRIMARY KEY (user_id, category)
);

CREATE TABLE password_reset_tokens (
  id         UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id    UUID        NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  token      TEXT        NOT NULL UNIQUE,
  expires_at TIMESTAMPTZ NOT NULL,
  used_at    TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_prt_token   ON password_reset_tokens (token) WHERE used_at IS NULL;
CREATE INDEX idx_prt_user_id ON password_reset_tokens (user_id);

CREATE TABLE tags (
  id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE courses (
  id             UUID           PRIMARY KEY DEFAULT gen_random_uuid(),
  title          TEXT           NOT NULL,
  slug           TEXT           NOT NULL UNIQUE,
  description    TEXT           NOT NULL,
  summary        TEXT           NOT NULL,
  category       TEXT           NOT NULL,
  level          course_level   NOT NULL,
  list_price     NUMERIC(10,2)  NOT NULL DEFAULT 0 CHECK (list_price >= 0),
  sale_price     NUMERIC(10,2)           CHECK (sale_price >= 0 AND sale_price < list_price),
  rating         NUMERIC(3,2)   NOT NULL DEFAULT 0 CHECK (rating BETWEEN 0 AND 5),
  student_count  INTEGER        NOT NULL DEFAULT 0 CHECK (student_count >= 0),
  instructor     TEXT           NOT NULL,
  instructor_bio TEXT           NOT NULL,
  cover          TEXT           NOT NULL,
  status         content_status NOT NULL DEFAULT 'draft',
  created_at     TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
  updated_at     TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
  fts_vector TSVECTOR GENERATED ALWAYS AS (
    setweight(to_tsvector('english', coalesce(title,       '')), 'A') ||
    setweight(to_tsvector('english', coalesce(summary,     '')), 'B') ||
    setweight(to_tsvector('english', coalesce(description, '')), 'C') ||
    setweight(to_tsvector('english', coalesce(instructor,  '')), 'D')
  ) STORED
);

CREATE INDEX idx_courses_slug       ON courses (slug);
CREATE INDEX idx_courses_category   ON courses (category);
CREATE INDEX idx_courses_level      ON courses (level);
CREATE INDEX idx_courses_status     ON courses (status);
CREATE INDEX idx_courses_list_price ON courses (list_price);
CREATE INDEX idx_courses_rating     ON courses (rating DESC);
CREATE INDEX idx_courses_fts        ON courses USING GIN (fts_vector);
CREATE INDEX idx_courses_title_trgm ON courses USING GIN (title gin_trgm_ops);

CREATE TRIGGER trg_courses_updated_at
  BEFORE UPDATE ON courses
  FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TABLE course_tags (
  course_id UUID NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
  tag_id    UUID NOT NULL REFERENCES tags    (id) ON DELETE CASCADE,
  PRIMARY KEY (course_id, tag_id)
);

CREATE TABLE lessons (
  id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  course_id   UUID        NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
  title       TEXT        NOT NULL,
  duration    TEXT        NOT NULL,
  video_url   TEXT,
  is_preview  BOOLEAN     NOT NULL DEFAULT FALSE,
  description TEXT,
  sort_order  SMALLINT    NOT NULL DEFAULT 0,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_lessons_course_id ON lessons (course_id, sort_order);

CREATE TABLE lesson_resources (
  id         UUID     PRIMARY KEY DEFAULT gen_random_uuid(),
  lesson_id  UUID     NOT NULL REFERENCES lessons (id) ON DELETE CASCADE,
  label      TEXT     NOT NULL,
  url        TEXT     NOT NULL,
  sort_order SMALLINT NOT NULL DEFAULT 0
);

CREATE INDEX idx_lesson_resources_lesson_id ON lesson_resources (lesson_id, sort_order);

CREATE TABLE reviews (
  id            UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  course_id     UUID        NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
  user_id       UUID                    REFERENCES users   (id) ON DELETE SET NULL,
  reviewer_name TEXT        NOT NULL,
  rating        SMALLINT    NOT NULL CHECK (rating BETWEEN 1 AND 5),
  comment       TEXT        NOT NULL,
  reviewed_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_reviews_course_id ON reviews (course_id, reviewed_at DESC);
CREATE INDEX idx_reviews_user_id   ON reviews (user_id);

CREATE TABLE quizzes (
  id            UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  course_id     UUID        NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
  title         TEXT        NOT NULL,
  description   TEXT,
  passing_score SMALLINT    NOT NULL DEFAULT 70 CHECK (passing_score BETWEEN 0 AND 100),
  created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_quizzes_course_id ON quizzes (course_id);

CREATE TABLE quiz_questions (
  id                  UUID     PRIMARY KEY DEFAULT gen_random_uuid(),
  quiz_id             UUID     NOT NULL REFERENCES quizzes (id) ON DELETE CASCADE,
  prompt              TEXT     NOT NULL,
  correct_option_keys TEXT[]   NOT NULL DEFAULT '{}',
  is_multi_select     BOOLEAN  NOT NULL DEFAULT FALSE,
  explanation         TEXT,
  sort_order          SMALLINT NOT NULL DEFAULT 0
);

CREATE INDEX idx_quiz_questions_quiz_id ON quiz_questions (quiz_id, sort_order);

CREATE TABLE quiz_options (
  id          UUID     PRIMARY KEY DEFAULT gen_random_uuid(),
  question_id UUID     NOT NULL REFERENCES quiz_questions (id) ON DELETE CASCADE,
  option_key  TEXT     NOT NULL,
  text        TEXT     NOT NULL,
  sort_order  SMALLINT NOT NULL DEFAULT 0,
  UNIQUE (question_id, option_key)
);

CREATE INDEX idx_quiz_options_question_id ON quiz_options (question_id, sort_order);

CREATE TABLE quiz_attempts (
  id           UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id      UUID        NOT NULL REFERENCES users   (id) ON DELETE CASCADE,
  quiz_id      UUID        NOT NULL REFERENCES quizzes (id) ON DELETE CASCADE,
  score        SMALLINT    NOT NULL CHECK (score BETWEEN 0 AND 100),
  passed       BOOLEAN     NOT NULL,
  answers      JSONB       NOT NULL DEFAULT '{}',
  attempted_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_quiz_attempts_user_quiz ON quiz_attempts (user_id, quiz_id, attempted_at DESC);

CREATE TABLE enrollments (
  id                 UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id            UUID        NOT NULL REFERENCES users   (id) ON DELETE CASCADE,
  course_id          UUID        NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
  passed_quiz_ids    TEXT[]      NOT NULL DEFAULT '{}',
  time_spent_seconds INTEGER     NOT NULL DEFAULT 0,
  enrolled_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE (user_id, course_id)
);

CREATE INDEX idx_enrollments_user_id   ON enrollments (user_id);
CREATE INDEX idx_enrollments_course_id ON enrollments (course_id);

CREATE TRIGGER trg_enrollments_updated_at
  BEFORE UPDATE ON enrollments
  FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TABLE enrollment_lesson_progress (
  enrollment_id UUID        NOT NULL REFERENCES enrollments (id) ON DELETE CASCADE,
  lesson_id     UUID        NOT NULL REFERENCES lessons     (id) ON DELETE CASCADE,
  time_spent_s  INTEGER     NOT NULL DEFAULT 0 CHECK (time_spent_s >= 0),
  completed_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (enrollment_id, lesson_id)
);

CREATE INDEX idx_elp_enrollment_id ON enrollment_lesson_progress (enrollment_id);
CREATE INDEX idx_elp_lesson_id     ON enrollment_lesson_progress (lesson_id);

CREATE TABLE transactions (
  id         UUID               PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id    UUID                           REFERENCES users   (id) ON DELETE SET NULL,
  course_id  UUID                           REFERENCES courses (id) ON DELETE SET NULL,
  amount     NUMERIC(10,2)      NOT NULL CHECK (amount >= 0),
  status     transaction_status NOT NULL,
  created_at TIMESTAMPTZ        NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_transactions_user_id    ON transactions (user_id);
CREATE INDEX idx_transactions_course_id  ON transactions (course_id);
CREATE INDEX idx_transactions_status     ON transactions (status);
CREATE INDEX idx_transactions_created_at ON transactions (created_at DESC);

CREATE TABLE wishlists (
  user_id   UUID        NOT NULL REFERENCES users   (id) ON DELETE CASCADE,
  course_id UUID        NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
  added_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (user_id, course_id)
);

CREATE INDEX idx_wishlists_user_id ON wishlists (user_id);

CREATE TABLE chat_sessions (
  id         UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id    UUID        NOT NULL REFERENCES users   (id) ON DELETE CASCADE,
  course_id  UUID                    REFERENCES courses (id) ON DELETE SET NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE NULLS NOT DISTINCT (user_id, course_id)
);

CREATE INDEX idx_chat_sessions_user_id ON chat_sessions (user_id);

CREATE TRIGGER trg_chat_sessions_updated_at
  BEFORE UPDATE ON chat_sessions
  FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TABLE chat_messages (
  id         UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  session_id UUID        NOT NULL REFERENCES chat_sessions (id) ON DELETE CASCADE,
  role       chat_role   NOT NULL,
  content    TEXT        NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_chat_messages_session_created ON chat_messages (session_id, created_at ASC);

CREATE TABLE certificates (
  id             UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id        UUID        NOT NULL REFERENCES users       (id) ON DELETE CASCADE,
  course_id      UUID        NOT NULL REFERENCES courses     (id) ON DELETE CASCADE,
  enrollment_id  UUID        NOT NULL REFERENCES enrollments (id) ON DELETE CASCADE UNIQUE,
  issued_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  download_count INTEGER     NOT NULL DEFAULT 0 CHECK (download_count >= 0)
);

CREATE INDEX idx_certificates_user_id   ON certificates (user_id);
CREATE INDEX idx_certificates_course_id ON certificates (course_id);

CREATE TABLE notifications (
  id         UUID              PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id    UUID              NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  type       notification_type NOT NULL,
  title      TEXT              NOT NULL,
  body       TEXT,
  is_read    BOOLEAN           NOT NULL DEFAULT FALSE,
  related_id UUID,
  created_at TIMESTAMPTZ       NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_notifications_user_unread ON notifications (user_id, is_read, created_at DESC);

CREATE TABLE audit_log (
  id         BIGSERIAL   PRIMARY KEY,
  actor_id   UUID                    REFERENCES users (id) ON DELETE SET NULL,
  action     TEXT        NOT NULL,
  table_name TEXT        NOT NULL,
  record_id  TEXT        NOT NULL,
  old_data   JSONB,
  new_data   JSONB,
  ip_address INET,
  user_agent TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_audit_log_actor_id    ON audit_log (actor_id);
CREATE INDEX idx_audit_log_table_name  ON audit_log (table_name, record_id);
CREATE INDEX idx_audit_log_created_at  ON audit_log (created_at DESC);

CREATE TABLE promo_codes (
  id             UUID          PRIMARY KEY DEFAULT gen_random_uuid(),
  code           TEXT          NOT NULL UNIQUE,
  discount_type  discount_type NOT NULL DEFAULT 'percent',
  discount_value NUMERIC(10,2) NOT NULL CHECK (discount_value > 0),
  max_uses       INTEGER,
  use_count      INTEGER       NOT NULL DEFAULT 0 CHECK (use_count >= 0),
  expires_at     TIMESTAMPTZ,
  is_active      BOOLEAN       NOT NULL DEFAULT TRUE,
  created_at     TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_promo_codes_code ON promo_codes (code) WHERE is_active;

CREATE TABLE live_sessions (
  id               UUID           PRIMARY KEY DEFAULT gen_random_uuid(),
  title            TEXT           NOT NULL,
  instructor       TEXT           NOT NULL,
  description      TEXT           NOT NULL DEFAULT '',
  scheduled_at     TIMESTAMPTZ    NOT NULL,
  duration_minutes INTEGER        NOT NULL CHECK (duration_minutes > 0),
  topic            TEXT           NOT NULL DEFAULT 'All Topics',
  join_url         TEXT           NOT NULL,
  course_id        UUID                      REFERENCES courses (id) ON DELETE SET NULL,
  registered_count INTEGER        NOT NULL DEFAULT 0 CHECK (registered_count >= 0),
  max_attendees    INTEGER        NOT NULL CHECK (max_attendees > 0),
  list_price       NUMERIC(10,2)             CHECK (list_price >= 0),
  sale_price       NUMERIC(10,2)             CHECK (sale_price >= 0),
  status           content_status NOT NULL DEFAULT 'published',
  created_at       TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
  updated_at       TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_live_sessions_scheduled_at ON live_sessions (scheduled_at);
CREATE INDEX idx_live_sessions_status       ON live_sessions (status);
CREATE INDEX idx_live_sessions_course_id    ON live_sessions (course_id);

CREATE TRIGGER trg_live_sessions_updated_at
  BEFORE UPDATE ON live_sessions
  FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TABLE feature_flags (
  id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  name        TEXT        NOT NULL UNIQUE,
  label       TEXT        NOT NULL,
  description TEXT        NOT NULL DEFAULT '',
  enabled     BOOLEAN     NOT NULL DEFAULT TRUE,
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER trg_feature_flags_updated_at
  BEFORE UPDATE ON feature_flags
  FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- Views
CREATE OR REPLACE VIEW v_course_stats AS
SELECT
  c.id, c.title, c.slug, c.category, c.level, c.list_price, c.sale_price,
  c.status, c.instructor, c.cover, c.created_at,
  COUNT(DISTINCT e.id)                                                    AS enrollment_count,
  COUNT(DISTINCT r.id)                                                    AS review_count,
  ROUND(COALESCE(AVG(r.rating), 0), 2)                                   AS avg_rating,
  COALESCE(SUM(t.amount) FILTER (WHERE t.status = 'success'), 0)         AS revenue
FROM      courses      c
LEFT JOIN enrollments  e ON e.course_id = c.id
LEFT JOIN reviews      r ON r.course_id = c.id
LEFT JOIN transactions t ON t.course_id = c.id
GROUP BY c.id;

CREATE OR REPLACE VIEW v_enrollment_progress AS
SELECT
  e.id                                                                    AS enrollment_id,
  e.user_id,
  e.course_id,
  e.enrolled_at,
  COUNT(l.id)                                                             AS total_lessons,
  COUNT(elp.lesson_id)                                                    AS completed_lessons,
  CASE
    WHEN COUNT(l.id) = 0 THEN 0
    ELSE ROUND(COUNT(elp.lesson_id) * 100.0 / COUNT(l.id))
  END                                                                     AS progress_pct,
  (COUNT(l.id) > 0 AND COUNT(elp.lesson_id) = COUNT(l.id))              AS is_completed
FROM      enrollments                e
JOIN      lessons                    l   ON  l.course_id    = e.course_id
LEFT JOIN enrollment_lesson_progress elp ON  elp.enrollment_id = e.id
                                         AND elp.lesson_id     = l.id
GROUP BY e.id;

CREATE OR REPLACE VIEW v_admin_overview AS
SELECT
  (SELECT COUNT(*)             FROM users        WHERE role = 'user')          AS total_users,
  (SELECT COUNT(*)             FROM courses      WHERE status = 'published')   AS total_courses,
  (SELECT COUNT(*)             FROM enrollments)                                AS total_enrollments,
  (SELECT COALESCE(SUM(amount), 0)
                               FROM transactions WHERE status = 'success')     AS revenue;

CREATE OR REPLACE VIEW v_live_sessions AS
SELECT
  ls.*,
  CASE
    WHEN NOW() < ls.scheduled_at THEN 'upcoming'::session_status
    WHEN NOW() < ls.scheduled_at + (ls.duration_minutes * INTERVAL '1 minute') THEN 'live'::session_status
    ELSE 'ended'::session_status
  END AS computed_status
FROM live_sessions ls;
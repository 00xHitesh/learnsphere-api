-- Seed data matching src/lib/seed.ts

INSERT INTO currencies (code, symbol, name, rate_vs_usd) VALUES
  ('USD', '$',  'US Dollar',          1),
  ('EUR', '€',  'Euro',               0.92),
  ('GBP', '£',  'British Pound',      0.79),
  ('INR', '₹',  'Indian Rupee',       83.5),
  ('JPY', '¥',  'Japanese Yen',       149.5),
  ('AUD', 'A$', 'Australian Dollar',  1.52),
  ('CAD', 'C$', 'Canadian Dollar',    1.36),
  ('SGD', 'S$', 'Singapore Dollar',   1.34),
  ('AED', 'د.إ','UAE Dirham',         3.67),
  ('BRL', 'R$', 'Brazilian Real',     4.97),
  ('MXN', '$',  'Mexican Peso',       17.2),
  ('CHF', 'Fr', 'Swiss Franc',        0.9),
  ('ZAR', 'R',  'South African Rand', 18.6),
  ('KRW', '₩',  'South Korean Won',  1325);

INSERT INTO regions (code, name, default_currency_code) VALUES
  ('US', 'United States',        'USD'),
  ('GB', 'United Kingdom',       'GBP'),
  ('IN', 'India',                'INR'),
  ('AU', 'Australia',            'AUD'),
  ('CA', 'Canada',               'CAD'),
  ('DE', 'Germany',              'EUR'),
  ('FR', 'France',               'EUR'),
  ('IT', 'Italy',                'EUR'),
  ('ES', 'Spain',                'EUR'),
  ('NL', 'Netherlands',          'EUR'),
  ('JP', 'Japan',                'JPY'),
  ('SG', 'Singapore',            'SGD'),
  ('AE', 'United Arab Emirates', 'AED'),
  ('BR', 'Brazil',               'BRL'),
  ('MX', 'Mexico',               'MXN'),
  ('CH', 'Switzerland',          'CHF'),
  ('ZA', 'South Africa',         'ZAR'),
  ('KR', 'South Korea',          'KRW'),
  ('NZ', 'New Zealand',          'AUD'),
  ('SE', 'Sweden',               'EUR'),
  ('NO', 'Norway',               'EUR'),
  ('DK', 'Denmark',              'EUR'),
  ('PL', 'Poland',               'EUR'),
  ('PT', 'Portugal',             'EUR'),
  ('BE', 'Belgium',              'EUR'),
  ('AT', 'Austria',              'EUR'),
  ('IE', 'Ireland',              'EUR'),
  ('FI', 'Finland',              'EUR'),
  ('HK', 'Hong Kong',            'USD'),
  ('MY', 'Malaysia',             'USD');

-- Passwords are placeholders — replace with real bcrypt hashes before production use
INSERT INTO users (id, name, email, password_hash, role) VALUES
  ('00000000-0000-0000-0000-000000000001', 'Admin',        'admin@lumen.io', '$2b$12$placeholder_hash_admin_CHANGEME', 'admin'),
  ('00000000-0000-0000-0000-000000000002', 'Demo Student', 'demo@lumen.io',  '$2b$12$placeholder_hash_demo_CHANGEME',  'user');

INSERT INTO user_interests (user_id, category) VALUES
  ('00000000-0000-0000-0000-000000000002', 'Design'),
  ('00000000-0000-0000-0000-000000000002', 'Programming');

INSERT INTO tags (id, name) VALUES
  ('aa000001-0000-0000-0000-000000000000', 'typography'),
  ('aa000002-0000-0000-0000-000000000000', 'design'),
  ('aa000003-0000-0000-0000-000000000000', 'layout'),
  ('aa000004-0000-0000-0000-000000000000', 'typescript'),
  ('aa000005-0000-0000-0000-000000000000', 'programming'),
  ('aa000006-0000-0000-0000-000000000000', 'javascript'),
  ('aa000007-0000-0000-0000-000000000000', 'art'),
  ('aa000008-0000-0000-0000-000000000000', 'watercolor'),
  ('aa000009-0000-0000-0000-000000000000', 'painting'),
  ('aa000010-0000-0000-0000-000000000000', 'product'),
  ('aa000011-0000-0000-0000-000000000000', 'strategy'),
  ('aa000012-0000-0000-0000-000000000000', 'business'),
  ('aa000013-0000-0000-0000-000000000000', 'react'),
  ('aa000014-0000-0000-0000-000000000000', 'architecture'),
  ('aa000015-0000-0000-0000-000000000000', 'frontend'),
  ('aa000016-0000-0000-0000-000000000000', 'writing'),
  ('aa000017-0000-0000-0000-000000000000', 'content');

INSERT INTO courses (id, title, slug, description, summary, category, level, list_price, sale_price, rating, student_count, instructor, instructor_bio, cover, status, created_at) VALUES
  ('cc000001-0000-0000-0000-000000000000', 'The Craft of Modern Typography',   'craft-of-modern-typography',   'A studio course on letterforms, hierarchy, and the quiet decisions that make pages feel composed rather than decorated.', 'Learn the principles behind elegant type systems, from kerning to scale, with practical studio exercises.', 'Design',      'Intermediate', 0,   NULL, 4.8, 2412, 'Mira Hollander', 'Type designer, formerly at Commercial Type. Teaches at RISD.',      'Aa', 'published', '2025-01-04 00:00:00+00'),
  ('cc000002-0000-0000-0000-000000000000', 'Foundations of TypeScript',        'foundations-of-typescript',    'Move from JavaScript to confident TypeScript: types, generics, narrowing, and patterns that scale across teams.',         'Master TypeScript fundamentals with practical patterns for real codebases.',                              'Programming', 'Beginner',     49,  29,   4.7, 8810, 'Kenji Romero',   'Senior engineer, contributor to several open-source TS projects.',  'TS', 'published', '2025-01-12 00:00:00+00'),
  ('cc000003-0000-0000-0000-000000000000', 'Watercolor: A Slow Practice',      'watercolor-slow-practice',     'An unhurried approach to watercolor — washes, restraint, and learning to leave the page alone.',                         'A meditative watercolor course focusing on restraint, washes, and quiet composition.',                    'Art',         'Beginner',     39,  24,   4.9, 1320, 'Eloise Marchetti','Painter and illustrator based in Florence.',                       '✦', 'published', '2025-02-02 00:00:00+00'),
  ('cc000004-0000-0000-0000-000000000000', 'Product Strategy for Builders',    'product-strategy-for-builders','How to choose what to build, how to know when you''re wrong, and how to write down a strategy people can actually follow.','Practical product strategy frameworks for technical founders and PMs.',                                   'Business',    'Intermediate', 79,  NULL, 4.6, 5230, 'Harper Liang',   'Former product lead at three early-stage startups.',                '◐', 'published', '2025-02-09 00:00:00+00'),
  ('cc000005-0000-0000-0000-000000000000', 'The Architecture of React Apps',   'architecture-of-react-apps',   'Composition, data flow, server boundaries, and the small choices that determine whether a codebase ages well.',          'Architectural principles for React applications that scale.',                                             'Programming', 'Advanced',     129, 89,   4.9, 3104, 'Sasha Vance',    'Staff engineer working on design systems.',                         '⚛', 'published', '2025-02-20 00:00:00+00'),
  ('cc000006-0000-0000-0000-000000000000', 'Writing for the Web',              'writing-for-the-web',          'Plain prose, useful headlines, and how to make a page feel like one mind wrote it. For designers, founders, and engineers.','Practical writing course for product, marketing, and documentation copy.',                                'Writing',     'Beginner',     0,   NULL, 4.7, 4112, 'Dale Okonkwo',   'Editor and writing coach.',                                         '¶', 'published', '2025-03-01 00:00:00+00');

INSERT INTO course_tags (course_id, tag_id) VALUES
  ('cc000001-0000-0000-0000-000000000000', 'aa000001-0000-0000-0000-000000000000'),
  ('cc000001-0000-0000-0000-000000000000', 'aa000002-0000-0000-0000-000000000000'),
  ('cc000001-0000-0000-0000-000000000000', 'aa000003-0000-0000-0000-000000000000'),
  ('cc000002-0000-0000-0000-000000000000', 'aa000004-0000-0000-0000-000000000000'),
  ('cc000002-0000-0000-0000-000000000000', 'aa000005-0000-0000-0000-000000000000'),
  ('cc000002-0000-0000-0000-000000000000', 'aa000006-0000-0000-0000-000000000000'),
  ('cc000003-0000-0000-0000-000000000000', 'aa000007-0000-0000-0000-000000000000'),
  ('cc000003-0000-0000-0000-000000000000', 'aa000008-0000-0000-0000-000000000000'),
  ('cc000003-0000-0000-0000-000000000000', 'aa000009-0000-0000-0000-000000000000'),
  ('cc000004-0000-0000-0000-000000000000', 'aa000010-0000-0000-0000-000000000000'),
  ('cc000004-0000-0000-0000-000000000000', 'aa000011-0000-0000-0000-000000000000'),
  ('cc000004-0000-0000-0000-000000000000', 'aa000012-0000-0000-0000-000000000000'),
  ('cc000005-0000-0000-0000-000000000000', 'aa000013-0000-0000-0000-000000000000'),
  ('cc000005-0000-0000-0000-000000000000', 'aa000014-0000-0000-0000-000000000000'),
  ('cc000005-0000-0000-0000-000000000000', 'aa000015-0000-0000-0000-000000000000'),
  ('cc000006-0000-0000-0000-000000000000', 'aa000016-0000-0000-0000-000000000000'),
  ('cc000006-0000-0000-0000-000000000000', 'aa000017-0000-0000-0000-000000000000');

INSERT INTO feature_flags (name, label, description, enabled) VALUES
  ('courses',       'Courses',       'Enable the courses catalog and all course detail pages.', TRUE),
  ('live_sessions', 'Live Sessions', 'Enable the live sessions section and the /live page.',    TRUE);

INSERT INTO promo_codes (code, discount_type, discount_value) VALUES
  ('WELCOME',  'percent', 10),
  ('SUMMER20', 'percent', 20),
  ('STUDY50',  'percent', 50);

INSERT INTO live_sessions (title, instructor, description, scheduled_at, duration_minutes, topic, join_url, registered_count, max_attendees) VALUES
  ('AI & Learning — Open Office Hours',    'The Lumen Team',   'Bring any question about courses, AI tools, or learning strategies. No agenda, all welcome.',              '2026-05-04 12:00:00+00', 2880, 'All Topics',  'https://meet.example.com/open-hours',        312, 500),
  ('Design Thinking Masterclass',          'Elena Vasquez',    'Explore design thinking frameworks with live exercises and a hands-on critique session.',                  '2026-05-07 15:00:00+00', 90,   'Design',      'https://meet.example.com/design-masterclass', 127, 200),
  ('Python Bootcamp — Office Hours',       'Marcus Chen',      'Got questions from this week''s Python lessons? Bring them here. Live coding, live answers.',              '2026-05-09 17:00:00+00', 60,   'Programming', 'https://meet.example.com/python-office-hours', 84, 150),
  ('Business Strategy Workshop',           'Aisha Okonkwo',    'Build a 1-page business strategy using real-world frameworks. Collaborative and fast-paced.',              '2026-05-12 14:00:00+00', 120,  'Business',    'https://meet.example.com/strategy-workshop',   56, 100),
  ('Watercolor for Beginners — Live Demo', 'Sofia Reyes',      'Watch and paint along as Sofia walks you through your first watercolor piece, start to finish.',           '2026-05-14 16:00:00+00', 75,   'Art',         'https://meet.example.com/watercolor-demo',    203, 300);
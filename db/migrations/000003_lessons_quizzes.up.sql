-- UUID key:
--   bb = lessons  (bb000001–bb000030)
--   dd = quizzes  (dd000001–dd000006)
--   ee = questions(ee000001–ee000024)
--   ff = options  (ff000001–ff000096)

-- ============================================================
-- LESSONS  (5 per course)
-- ============================================================
INSERT INTO lessons (id, course_id, title, duration, is_preview, description, sort_order) VALUES

  -- The Craft of Modern Typography
  ('bb000001-0000-0000-0000-000000000000', 'cc000001-0000-0000-0000-000000000000',
   'Introduction to Letterforms',   '12:00', TRUE,
   'An overview of type anatomy: stems, bowls, ascenders, descenders, and the invisible grid beneath every typeface.', 1),
  ('bb000002-0000-0000-0000-000000000000', 'cc000001-0000-0000-0000-000000000000',
   'Hierarchy and Scale',            '18:30', FALSE,
   'How size, weight, and contrast direct the reader''s eye — build a three-level hierarchy from scratch.', 2),
  ('bb000003-0000-0000-0000-000000000000', 'cc000001-0000-0000-0000-000000000000',
   'Kerning and Spacing',            '22:15', FALSE,
   'The subtle art of adjusting space between individual letter pairs and the optical corrections that make type feel right.', 3),
  ('bb000004-0000-0000-0000-000000000000', 'cc000001-0000-0000-0000-000000000000',
   'Type Systems and Families',      '25:00', FALSE,
   'Pairing typefaces and building a consistent type system — fonts that work together without competing.', 4),
  ('bb000005-0000-0000-0000-000000000000', 'cc000001-0000-0000-0000-000000000000',
   'Layout and Composition',         '28:45', FALSE,
   'Applying typographic principles to full page layouts: grid, white space, and the quiet decisions that make pages feel composed.', 5),

  -- Foundations of TypeScript
  ('bb000006-0000-0000-0000-000000000000', 'cc000002-0000-0000-0000-000000000000',
   'Why TypeScript?',                '10:00', TRUE,
   'The case for static types: catch bugs earlier, document intent, and scale confidently across large codebases.', 1),
  ('bb000007-0000-0000-0000-000000000000', 'cc000002-0000-0000-0000-000000000000',
   'Basic Types and Inference',      '20:00', FALSE,
   'Primitives, arrays, tuples, and how the compiler infers types so you don''t have to annotate everything.', 2),
  ('bb000008-0000-0000-0000-000000000000', 'cc000002-0000-0000-0000-000000000000',
   'Interfaces and Type Aliases',    '24:30', FALSE,
   'Defining object shapes, extending interfaces, and knowing when to reach for a type alias instead.', 3),
  ('bb000009-0000-0000-0000-000000000000', 'cc000002-0000-0000-0000-000000000000',
   'Generics',                       '30:00', FALSE,
   'Writing reusable, type-safe functions and data structures with generics. Constraints and conditional types included.', 4),
  ('bb000010-0000-0000-0000-000000000000', 'cc000002-0000-0000-0000-000000000000',
   'Narrowing and Type Guards',      '26:15', FALSE,
   'Narrowing union types with typeof, instanceof, discriminated unions, and custom type guard functions.', 5),

  -- Watercolor: A Slow Practice
  ('bb000011-0000-0000-0000-000000000000', 'cc000003-0000-0000-0000-000000000000',
   'Materials and Setup',            '14:00', TRUE,
   'Choosing paper, brushes, and paints. Why cheap materials slow your learning and what''s worth investing in.', 1),
  ('bb000012-0000-0000-0000-000000000000', 'cc000003-0000-0000-0000-000000000000',
   'Your First Washes',              '20:00', FALSE,
   'Flat and graded washes from scratch. Loading the brush, controlling water ratio, and reading wet paper.', 2),
  ('bb000013-0000-0000-0000-000000000000', 'cc000003-0000-0000-0000-000000000000',
   'Wet-on-Wet Technique',           '22:30', FALSE,
   'Painting into a wet surface to create soft, diffused edges. Timing and humidity are everything.', 3),
  ('bb000014-0000-0000-0000-000000000000', 'cc000003-0000-0000-0000-000000000000',
   'Dry Brush and Texture',          '18:45', FALSE,
   'Using a nearly dry brush on rough paper for texture and detail. When to embrace the happy accident.', 4),
  ('bb000015-0000-0000-0000-000000000000', 'cc000003-0000-0000-0000-000000000000',
   'Composition and Restraint',      '24:00', FALSE,
   'The hardest lesson: knowing when to stop. Negative space, the rule of thirds, and editing your instincts.', 5),

  -- Product Strategy for Builders
  ('bb000016-0000-0000-0000-000000000000', 'cc000004-0000-0000-0000-000000000000',
   'What is Strategy?',              '15:00', TRUE,
   'Strategy is not a roadmap. It''s a coherent set of choices that explain why you''re doing what you''re doing — and what you''re not doing.', 1),
  ('bb000017-0000-0000-0000-000000000000', 'cc000004-0000-0000-0000-000000000000',
   'Identifying the Right Problem',  '22:00', FALSE,
   'How to avoid building the wrong thing well. Jobs-to-be-done, user interviews, and separating symptoms from root causes.', 2),
  ('bb000018-0000-0000-0000-000000000000', 'cc000004-0000-0000-0000-000000000000',
   'Writing a Strategy Document',    '28:00', FALSE,
   'A one-page strategy template — situation, complication, resolution — written in plain language your team can act on.', 3),
  ('bb000019-0000-0000-0000-000000000000', 'cc000004-0000-0000-0000-000000000000',
   'Validating Assumptions',         '20:30', FALSE,
   'How to know you''re wrong early. Assumption mapping, cheap experiments, and what actually counts as signal.', 4),
  ('bb000020-0000-0000-0000-000000000000', 'cc000004-0000-0000-0000-000000000000',
   'Communicating to Stakeholders',  '24:45', FALSE,
   'Presenting strategy to people who didn''t make it. Narratives, trade-off framing, and alignment meetings that work.', 5),

  -- The Architecture of React Apps
  ('bb000021-0000-0000-0000-000000000000', 'cc000005-0000-0000-0000-000000000000',
   'Component Composition Patterns', '18:00', TRUE,
   'Slots, compound components, and render props — the mental model behind each and when each is the right tool.', 1),
  ('bb000022-0000-0000-0000-000000000000', 'cc000005-0000-0000-0000-000000000000',
   'State Management Trade-offs',    '32:00', FALSE,
   'useState vs useReducer vs context vs Zustand vs server state. A decision framework, not a framework recommendation.', 2),
  ('bb000023-0000-0000-0000-000000000000', 'cc000005-0000-0000-0000-000000000000',
   'Data Flow and Server Boundaries','28:30', FALSE,
   'Server components, client components, and the boundary between them. Where data fetching actually belongs.', 3),
  ('bb000024-0000-0000-0000-000000000000', 'cc000005-0000-0000-0000-000000000000',
   'Code Splitting and Performance', '24:15', FALSE,
   'Lazy loading, Suspense boundaries, bundle analysis, and measuring what you actually improved.', 4),
  ('bb000025-0000-0000-0000-000000000000', 'cc000005-0000-0000-0000-000000000000',
   'Testing Architecture',           '26:00', FALSE,
   'Testing components in isolation vs integration. What to mock, what to hit for real, and the tests that catch regressions.', 5),

  -- Writing for the Web
  ('bb000026-0000-0000-0000-000000000000', 'cc000006-0000-0000-0000-000000000000',
   'Why Web Writing Is Different',   '12:00', TRUE,
   'Readers scan, not read. The F-pattern, cognitive load, and how good web writing works with the grain instead of against it.', 1),
  ('bb000027-0000-0000-0000-000000000000', 'cc000006-0000-0000-0000-000000000000',
   'Headlines That Work',            '18:30', FALSE,
   'What makes a headline useful vs clever. Specificity, front-loading, and avoiding the curiosity-gap trap.', 2),
  ('bb000028-0000-0000-0000-000000000000', 'cc000006-0000-0000-0000-000000000000',
   'Plain Prose Principles',         '22:00', FALSE,
   'Short sentences. Active voice. No throat-clearing. Six principles for writing that respects the reader''s time.', 3),
  ('bb000029-0000-0000-0000-000000000000', 'cc000006-0000-0000-0000-000000000000',
   'Structure and Scanability',      '20:15', FALSE,
   'Headers, bullets, and white space. How to organize a page so busy readers get what they need without reading everything.', 4),
  ('bb000030-0000-0000-0000-000000000000', 'cc000006-0000-0000-0000-000000000000',
   'Voice and Consistency',          '24:30', FALSE,
   'What a voice actually is, how to define it, and how to keep it consistent across a team of writers.', 5);


-- ============================================================
-- QUIZZES  (1 per course)
-- ============================================================
INSERT INTO quizzes (id, course_id, title, description, passing_score) VALUES
  ('dd000001-0000-0000-0000-000000000000', 'cc000001-0000-0000-0000-000000000000',
   'Typography Fundamentals',
   'Test your understanding of type anatomy, hierarchy, spacing, and layout.', 70),
  ('dd000002-0000-0000-0000-000000000000', 'cc000002-0000-0000-0000-000000000000',
   'TypeScript Basics',
   'Test your knowledge of TypeScript types, inference, generics, and narrowing.', 70),
  ('dd000003-0000-0000-0000-000000000000', 'cc000003-0000-0000-0000-000000000000',
   'Watercolor Essentials',
   'Check your grasp of materials, techniques, and composition principles.', 70),
  ('dd000004-0000-0000-0000-000000000000', 'cc000004-0000-0000-0000-000000000000',
   'Product Strategy Check',
   'Assess your understanding of strategy frameworks, problem framing, and validation.', 70),
  ('dd000005-0000-0000-0000-000000000000', 'cc000005-0000-0000-0000-000000000000',
   'React Architecture',
   'Test your knowledge of composition patterns, state management, and performance.', 70),
  ('dd000006-0000-0000-0000-000000000000', 'cc000006-0000-0000-0000-000000000000',
   'Web Writing Fundamentals',
   'Check your understanding of web writing principles, headlines, and structure.', 70);


-- ============================================================
-- QUIZ QUESTIONS  (4 per quiz)
-- ============================================================
INSERT INTO quiz_questions (id, quiz_id, prompt, correct_option_keys, is_multi_select, explanation, sort_order) VALUES

  -- Quiz 1: Typography Fundamentals
  ('ee000001-0000-0000-0000-000000000000', 'dd000001-0000-0000-0000-000000000000',
   'What is kerning?',
   ARRAY['b'], FALSE,
   'Kerning is the manual adjustment of space between a specific pair of characters to improve visual balance. Tracking adjusts all characters uniformly; leading adjusts line height.', 1),

  ('ee000002-0000-0000-0000-000000000000', 'dd000001-0000-0000-0000-000000000000',
   'In a typographic hierarchy, which element should carry the greatest visual weight?',
   ARRAY['c'], FALSE,
   'Headings establish the top level of the hierarchy through larger size, heavier weight, or stronger contrast — they are the first thing a reader notices.', 2),

  ('ee000003-0000-0000-0000-000000000000', 'dd000001-0000-0000-0000-000000000000',
   'What does "tracking" (letter-spacing) control?',
   ARRAY['c'], FALSE,
   'Tracking applies a uniform spacing adjustment across a selected range of characters. Kerning targets a specific pair; leading controls vertical line spacing.', 3),

  ('ee000004-0000-0000-0000-000000000000', 'dd000001-0000-0000-0000-000000000000',
   'Which of the following best describes a "type system"?',
   ARRAY['d'], FALSE,
   'A type system is a curated set of typeface choices, size scales, weights, and spacing rules that work together consistently — it is a design decision, not a software tool.', 4),

  -- Quiz 2: TypeScript Basics
  ('ee000005-0000-0000-0000-000000000000', 'dd000002-0000-0000-0000-000000000000',
   'What is the key difference between `unknown` and `any` in TypeScript?',
   ARRAY['b'], FALSE,
   'any opts out of type checking entirely. unknown preserves safety by requiring a type assertion or narrowing check before the value can be used in any typed operation.', 1),

  ('ee000006-0000-0000-0000-000000000000', 'dd000002-0000-0000-0000-000000000000',
   'What does TypeScript''s type inference do?',
   ARRAY['a'], FALSE,
   'The compiler analyses assigned values and usage context to deduce types automatically, meaning you only need to add explicit annotations where inference cannot determine the correct type.', 2),

  ('ee000007-0000-0000-0000-000000000000', 'dd000002-0000-0000-0000-000000000000',
   'Which syntax correctly defines a generic function in TypeScript?',
   ARRAY['b'], FALSE,
   'The type parameter is declared with angle brackets immediately before the parameter list: function identity<T>(arg: T): T. The other forms are not valid TypeScript syntax.', 3),

  ('ee000008-0000-0000-0000-000000000000', 'dd000002-0000-0000-0000-000000000000',
   'What is a discriminated union?',
   ARRAY['d'], FALSE,
   'A discriminated union is a union type where every member shares a common literal property (the discriminant). TypeScript uses that property to narrow to the correct variant in if/switch blocks.', 4),

  -- Quiz 3: Watercolor Essentials
  ('ee000009-0000-0000-0000-000000000000', 'dd000003-0000-0000-0000-000000000000',
   'What is the wet-on-wet technique?',
   ARRAY['b'], FALSE,
   'Wet-on-wet means applying wet pigment to paper that has already been wetted. This creates soft, diffused edges because the pigment spreads into the water already on the surface.', 1),

  ('ee000010-0000-0000-0000-000000000000', 'dd000003-0000-0000-0000-000000000000',
   'Which paper weight is most commonly recommended for watercolor painting?',
   ARRAY['c'], FALSE,
   '140 lb (300 gsm) is the standard recommendation. It is heavy enough to resist buckling through multiple wet washes without requiring pre-stretching in most situations.', 2),

  ('ee000011-0000-0000-0000-000000000000', 'dd000003-0000-0000-0000-000000000000',
   'Why does this course emphasise the principle of restraint?',
   ARRAY['c'], FALSE,
   'Watercolor''s defining quality is luminous transparency. Over-working a wet area muddies the pigment, destroys the transparency, and cannot be fully undone once dry.', 3),

  ('ee000012-0000-0000-0000-000000000000', 'dd000003-0000-0000-0000-000000000000',
   'What does the rule of thirds suggest in composition?',
   ARRAY['b'], FALSE,
   'Dividing the canvas into a 3×3 grid and placing subjects along the dividing lines or at their intersections creates more dynamic compositions than centring the subject.', 4),

  -- Quiz 4: Product Strategy Check
  ('ee000013-0000-0000-0000-000000000000', 'dd000004-0000-0000-0000-000000000000',
   'Which of the following best describes a product strategy?',
   ARRAY['a'], FALSE,
   'Strategy is a coherent set of choices — where to play, how to win, and what not to do. A roadmap is a delivery schedule; strategy explains the reasoning behind it.', 1),

  ('ee000014-0000-0000-0000-000000000000', 'dd000004-0000-0000-0000-000000000000',
   'What is the "Jobs-to-be-Done" framework primarily concerned with?',
   ARRAY['b'], FALSE,
   'JTBD focuses on the underlying progress a customer is trying to make. People "hire" products to do a job; understanding that job reveals what the product must actually do.', 2),

  ('ee000015-0000-0000-0000-000000000000', 'dd000004-0000-0000-0000-000000000000',
   'What is the goal of assumption mapping?',
   ARRAY['c'], FALSE,
   'Every strategy rests on beliefs that may be wrong. Assumption mapping surfaces the riskiest ones so you can design cheap, fast experiments to validate or invalidate them before over-investing.', 3),

  ('ee000016-0000-0000-0000-000000000000', 'dd000004-0000-0000-0000-000000000000',
   'When presenting strategy to stakeholders, what matters most?',
   ARRAY['d'], FALSE,
   'Stakeholders need to understand the situation you are responding to, the bet you are placing, and what you expect to happen — a clear narrative they can challenge or commit to.', 4),

  -- Quiz 5: React Architecture
  ('ee000017-0000-0000-0000-000000000000', 'dd000005-0000-0000-0000-000000000000',
   'What is the compound component pattern?',
   ARRAY['c'], FALSE,
   'Compound components expose a set of related child components. The parent holds shared state in context and the children consume it implicitly, keeping the public API clean.', 1),

  ('ee000018-0000-0000-0000-000000000000', 'dd000005-0000-0000-0000-000000000000',
   'When is React Context most appropriate for state management?',
   ARRAY['a'], FALSE,
   'Context is well-suited to low-frequency updates (theme, locale, auth user) shared across a subtree. Frequent updates cause every consumer to re-render, making it a poor fit for high-churn state.', 2),

  ('ee000019-0000-0000-0000-000000000000', 'dd000005-0000-0000-0000-000000000000',
   'What is a React Server Component?',
   ARRAY['a'], FALSE,
   'A Server Component runs exclusively on the server. Its output is sent as a serialized UI description — no component JavaScript is included in the client bundle, reducing bundle size significantly.', 3),

  ('ee000020-0000-0000-0000-000000000000', 'dd000005-0000-0000-0000-000000000000',
   'What is the primary purpose of React.lazy combined with Suspense?',
   ARRAY['c'], FALSE,
   'React.lazy wraps a dynamic import, splitting the bundle at that component boundary. Suspense shows a fallback while the chunk loads. This reduces initial load time by deferring non-critical code.', 4),

  -- Quiz 6: Web Writing Fundamentals
  ('ee000021-0000-0000-0000-000000000000', 'dd000006-0000-0000-0000-000000000000',
   'What is the "inverted pyramid" structure in web writing?',
   ARRAY['a'], FALSE,
   'The inverted pyramid leads with the most important information first, then adds supporting detail below. It respects readers who stop reading early and is the opposite of academic writing.', 1),

  ('ee000022-0000-0000-0000-000000000000', 'dd000006-0000-0000-0000-000000000000',
   'Which sentence is written in active voice?',
   ARRAY['c'], FALSE,
   'In active voice the subject performs the action: "The team submitted the report." Passive voice inverts this: "The report was submitted by the team" — wordier and often less clear.', 2),

  ('ee000023-0000-0000-0000-000000000000', 'dd000006-0000-0000-0000-000000000000',
   'What makes a web headline effective?',
   ARRAY['d'], FALSE,
   'Effective headlines are specific, front-load the key information, and make sense without reading the body. Clever or vague headlines force the reader to work too hard.', 3),

  ('ee000024-0000-0000-0000-000000000000', 'dd000006-0000-0000-0000-000000000000',
   'What does the F-pattern describe in the context of web reading?',
   ARRAY['b'], FALSE,
   'Eye-tracking studies show readers scan across the top of the page, then down the left side in a rough F shape. Good web structure puts the most important content along those axes.', 4);


-- ============================================================
-- QUIZ OPTIONS  (4 per question, a–d)
-- ============================================================
INSERT INTO quiz_options (id, question_id, option_key, text, sort_order) VALUES

  -- ee000001: What is kerning?
  ('ff000001-0000-0000-0000-000000000000', 'ee000001-0000-0000-0000-000000000000', 'a', 'The vertical space between lines of text', 1),
  ('ff000002-0000-0000-0000-000000000000', 'ee000001-0000-0000-0000-000000000000', 'b', 'The adjustment of space between a specific pair of adjacent characters', 2),
  ('ff000003-0000-0000-0000-000000000000', 'ee000001-0000-0000-0000-000000000000', 'c', 'A uniform spacing value applied across a range of characters', 3),
  ('ff000004-0000-0000-0000-000000000000', 'ee000001-0000-0000-0000-000000000000', 'd', 'The height of lowercase letters relative to capital letters', 4),

  -- ee000002: Greatest visual weight in hierarchy
  ('ff000005-0000-0000-0000-000000000000', 'ee000002-0000-0000-0000-000000000000', 'a', 'Body text, because there is more of it on the page', 1),
  ('ff000006-0000-0000-0000-000000000000', 'ee000002-0000-0000-0000-000000000000', 'b', 'Captions, because they sit closest to the images', 2),
  ('ff000007-0000-0000-0000-000000000000', 'ee000002-0000-0000-0000-000000000000', 'c', 'Headings, through larger size, heavier weight, or stronger contrast', 3),
  ('ff000008-0000-0000-0000-000000000000', 'ee000002-0000-0000-0000-000000000000', 'd', 'Footnotes, because their isolation draws the eye', 4),

  -- ee000003: What does tracking control?
  ('ff000009-0000-0000-0000-000000000000', 'ee000003-0000-0000-0000-000000000000', 'a', 'The space between two specific adjacent characters', 1),
  ('ff000010-0000-0000-0000-000000000000', 'ee000003-0000-0000-0000-000000000000', 'b', 'The vertical distance between text baselines', 2),
  ('ff000011-0000-0000-0000-000000000000', 'ee000003-0000-0000-0000-000000000000', 'c', 'The uniform letter-spacing applied across a range of characters', 3),
  ('ff000012-0000-0000-0000-000000000000', 'ee000003-0000-0000-0000-000000000000', 'd', 'The horizontal width of individual letterform strokes', 4),

  -- ee000004: What is a type system?
  ('ff000013-0000-0000-0000-000000000000', 'ee000004-0000-0000-0000-000000000000', 'a', 'A software application for rendering type on screen', 1),
  ('ff000014-0000-0000-0000-000000000000', 'ee000004-0000-0000-0000-000000000000', 'b', 'A single typeface that includes multiple weights and styles', 2),
  ('ff000015-0000-0000-0000-000000000000', 'ee000004-0000-0000-0000-000000000000', 'c', 'The baseline grid on which type is laid out', 3),
  ('ff000016-0000-0000-0000-000000000000', 'ee000004-0000-0000-0000-000000000000', 'd', 'A curated set of typefaces, scales, and spacing rules that work together consistently', 4),

  -- ee000005: unknown vs any
  ('ff000017-0000-0000-0000-000000000000', 'ee000005-0000-0000-0000-000000000000', 'a', 'unknown is only valid for object types; any works for all types', 1),
  ('ff000018-0000-0000-0000-000000000000', 'ee000005-0000-0000-0000-000000000000', 'b', 'any disables type checking on a value; unknown requires a type check before use', 2),
  ('ff000019-0000-0000-0000-000000000000', 'ee000005-0000-0000-0000-000000000000', 'c', 'They are identical; unknown is just a more readable alias for any', 3),
  ('ff000020-0000-0000-0000-000000000000', 'ee000005-0000-0000-0000-000000000000', 'd', 'unknown is a runtime type; any is a compile-time annotation only', 4),

  -- ee000006: Type inference
  ('ff000021-0000-0000-0000-000000000000', 'ee000006-0000-0000-0000-000000000000', 'a', 'The compiler automatically deduces types from assigned values and usage context', 1),
  ('ff000022-0000-0000-0000-000000000000', 'ee000006-0000-0000-0000-000000000000', 'b', 'It converts JavaScript files to TypeScript automatically', 2),
  ('ff000023-0000-0000-0000-000000000000', 'ee000006-0000-0000-0000-000000000000', 'c', 'It generates runtime type assertions in the compiled output', 3),
  ('ff000024-0000-0000-0000-000000000000', 'ee000006-0000-0000-0000-000000000000', 'd', 'It enables strict null checks across the entire codebase', 4),

  -- ee000007: Generic function syntax
  ('ff000025-0000-0000-0000-000000000000', 'ee000007-0000-0000-0000-000000000000', 'a', 'function identity(T)(arg: T): T', 1),
  ('ff000026-0000-0000-0000-000000000000', 'ee000007-0000-0000-0000-000000000000', 'b', 'function identity<T>(arg: T): T', 2),
  ('ff000027-0000-0000-0000-000000000000', 'ee000007-0000-0000-0000-000000000000', 'c', 'function identity(arg: generic): generic', 3),
  ('ff000028-0000-0000-0000-000000000000', 'ee000007-0000-0000-0000-000000000000', 'd', 'function<T> identity(arg: T): T', 4),

  -- ee000008: Discriminated union
  ('ff000029-0000-0000-0000-000000000000', 'ee000008-0000-0000-0000-000000000000', 'a', 'A union of two interfaces with no overlapping properties', 1),
  ('ff000030-0000-0000-0000-000000000000', 'ee000008-0000-0000-0000-000000000000', 'b', 'A union type that excludes null and undefined', 2),
  ('ff000031-0000-0000-0000-000000000000', 'ee000008-0000-0000-0000-000000000000', 'c', 'A union type restricted to use inside switch statements only', 3),
  ('ff000032-0000-0000-0000-000000000000', 'ee000008-0000-0000-0000-000000000000', 'd', 'A union type where all members share a common literal property used to narrow the type', 4),

  -- ee000009: Wet-on-wet
  ('ff000033-0000-0000-0000-000000000000', 'ee000009-0000-0000-0000-000000000000', 'a', 'Applying a dry brush over a layer of fully dried paint', 1),
  ('ff000034-0000-0000-0000-000000000000', 'ee000009-0000-0000-0000-000000000000', 'b', 'Painting wet pigment directly onto a pre-wetted paper surface', 2),
  ('ff000035-0000-0000-0000-000000000000', 'ee000009-0000-0000-0000-000000000000', 'c', 'Mixing paint with water on a palette before any contact with the paper', 3),
  ('ff000036-0000-0000-0000-000000000000', 'ee000009-0000-0000-0000-000000000000', 'd', 'Applying two quick strokes of wet paint on dry paper in rapid succession', 4),

  -- ee000010: Paper weight
  ('ff000037-0000-0000-0000-000000000000', 'ee000010-0000-0000-0000-000000000000', 'a', '60 lb (90 gsm) — standard sketching paper', 1),
  ('ff000038-0000-0000-0000-000000000000', 'ee000010-0000-0000-0000-000000000000', 'b', '90 lb (135 gsm) — light watercolor paper suitable for small studies', 2),
  ('ff000039-0000-0000-0000-000000000000', 'ee000010-0000-0000-0000-000000000000', 'c', '140 lb (300 gsm) — the standard weight recommended for most watercolor work', 3),
  ('ff000040-0000-0000-0000-000000000000', 'ee000010-0000-0000-0000-000000000000', 'd', '20 lb (75 gsm) — printer paper, useful for quick practice washes', 4),

  -- ee000011: Restraint in watercolor
  ('ff000041-0000-0000-0000-000000000000', 'ee000011-0000-0000-0000-000000000000', 'a', 'Watercolor dries darker, so you should always revisit and add more paint', 1),
  ('ff000042-0000-0000-0000-000000000000', 'ee000011-0000-0000-0000-000000000000', 'b', 'Painting too slowly causes the pigment to crack once it dries', 2),
  ('ff000043-0000-0000-0000-000000000000', 'ee000011-0000-0000-0000-000000000000', 'c', 'Over-working wet watercolor destroys the transparency that makes the medium distinctive', 3),
  ('ff000044-0000-0000-0000-000000000000', 'ee000011-0000-0000-0000-000000000000', 'd', 'Restraint matters only for beginners; experienced painters work freely', 4),

  -- ee000012: Rule of thirds
  ('ff000045-0000-0000-0000-000000000000', 'ee000012-0000-0000-0000-000000000000', 'a', 'Limit every painting to exactly three dominant colors', 1),
  ('ff000046-0000-0000-0000-000000000000', 'ee000012-0000-0000-0000-000000000000', 'b', 'Divide the canvas into nine equal parts; place key elements along the lines or at intersections', 2),
  ('ff000047-0000-0000-0000-000000000000', 'ee000012-0000-0000-0000-000000000000', 'c', 'Spend one third sketching, one third painting, and one third finishing each piece', 3),
  ('ff000048-0000-0000-0000-000000000000', 'ee000012-0000-0000-0000-000000000000', 'd', 'Work in no more than three separate wet sessions per painting', 4),

  -- ee000013: Product strategy definition
  ('ff000049-0000-0000-0000-000000000000', 'ee000013-0000-0000-0000-000000000000', 'a', 'A coherent set of choices about where to play and how to win', 1),
  ('ff000050-0000-0000-0000-000000000000', 'ee000013-0000-0000-0000-000000000000', 'b', 'A detailed feature roadmap with quarterly delivery milestones', 2),
  ('ff000051-0000-0000-0000-000000000000', 'ee000013-0000-0000-0000-000000000000', 'c', 'A prioritised list of user stories ranked by engineering effort', 3),
  ('ff000052-0000-0000-0000-000000000000', 'ee000013-0000-0000-0000-000000000000', 'd', 'A competitive landscape document used to support fundraising', 4),

  -- ee000014: Jobs-to-be-Done
  ('ff000053-0000-0000-0000-000000000000', 'ee000014-0000-0000-0000-000000000000', 'a', 'Assigning technical tasks to team members based on skill level', 1),
  ('ff000054-0000-0000-0000-000000000000', 'ee000014-0000-0000-0000-000000000000', 'b', 'Understanding the underlying goal a customer is trying to achieve when they choose a product', 2),
  ('ff000055-0000-0000-0000-000000000000', 'ee000014-0000-0000-0000-000000000000', 'c', 'Defining the engineering tasks required to ship each feature', 3),
  ('ff000056-0000-0000-0000-000000000000', 'ee000014-0000-0000-0000-000000000000', 'd', 'Measuring time users spend on each product feature in analytics', 4),

  -- ee000015: Assumption mapping
  ('ff000057-0000-0000-0000-000000000000', 'ee000015-0000-0000-0000-000000000000', 'a', 'Documenting every feature the product intends to ship eventually', 1),
  ('ff000058-0000-0000-0000-000000000000', 'ee000015-0000-0000-0000-000000000000', 'b', 'Mapping the user journey from first discovery through to activation', 2),
  ('ff000059-0000-0000-0000-000000000000', 'ee000015-0000-0000-0000-000000000000', 'c', 'Identifying and ranking the riskiest beliefs your strategy depends on so you can test them cheaply', 3),
  ('ff000060-0000-0000-0000-000000000000', 'ee000015-0000-0000-0000-000000000000', 'd', 'Creating a technical dependency map for the next release', 4),

  -- ee000016: Stakeholder communication
  ('ff000061-0000-0000-0000-000000000000', 'ee000016-0000-0000-0000-000000000000', 'a', 'Including as much supporting data as possible to justify every decision', 1),
  ('ff000062-0000-0000-0000-000000000000', 'ee000016-0000-0000-0000-000000000000', 'b', 'Getting explicit sign-off from every team before presenting to leadership', 2),
  ('ff000063-0000-0000-0000-000000000000', 'ee000016-0000-0000-0000-000000000000', 'c', 'Leading with competitive analysis to establish market urgency', 3),
  ('ff000064-0000-0000-0000-000000000000', 'ee000016-0000-0000-0000-000000000000', 'd', 'Presenting a clear narrative — the situation, the bet you''re making, and the expected outcome', 4),

  -- ee000017: Compound component pattern
  ('ff000065-0000-0000-0000-000000000000', 'ee000017-0000-0000-0000-000000000000', 'a', 'Using two separate React frameworks in the same application', 1),
  ('ff000066-0000-0000-0000-000000000000', 'ee000017-0000-0000-0000-000000000000', 'b', 'Splitting a single component into a server half and a client half', 2),
  ('ff000067-0000-0000-0000-000000000000', 'ee000017-0000-0000-0000-000000000000', 'c', 'A pattern where a parent component shares implicit state with its children through context', 3),
  ('ff000068-0000-0000-0000-000000000000', 'ee000017-0000-0000-0000-000000000000', 'd', 'A technique for composing multiple higher-order components in sequence', 4),

  -- ee000018: React Context
  ('ff000069-0000-0000-0000-000000000000', 'ee000018-0000-0000-0000-000000000000', 'a', 'For low-frequency state shared across a subtree, such as theme or locale', 1),
  ('ff000070-0000-0000-0000-000000000000', 'ee000018-0000-0000-0000-000000000000', 'b', 'As the single source of truth for all application state, replacing libraries like Redux', 2),
  ('ff000071-0000-0000-0000-000000000000', 'ee000018-0000-0000-0000-000000000000', 'c', 'For frequently updated server-fetched data to avoid prop drilling', 3),
  ('ff000072-0000-0000-0000-000000000000', 'ee000018-0000-0000-0000-000000000000', 'd', 'As the primary store for real-time data streams like WebSocket messages', 4),

  -- ee000019: React Server Component
  ('ff000073-0000-0000-0000-000000000000', 'ee000019-0000-0000-0000-000000000000', 'a', 'A component that runs on the server only and sends no JavaScript to the client', 1),
  ('ff000074-0000-0000-0000-000000000000', 'ee000019-0000-0000-0000-000000000000', 'b', 'Any component that fetches data from a REST or GraphQL API', 2),
  ('ff000075-0000-0000-0000-000000000000', 'ee000019-0000-0000-0000-000000000000', 'c', 'A component wrapped in React.lazy for deferred client-side loading', 3),
  ('ff000076-0000-0000-0000-000000000000', 'ee000019-0000-0000-0000-000000000000', 'd', 'A component rendered on the server and then rehydrated on the client', 4),

  -- ee000020: React.lazy + Suspense
  ('ff000077-0000-0000-0000-000000000000', 'ee000020-0000-0000-0000-000000000000', 'a', 'To defer rendering an entire page until all API responses have resolved', 1),
  ('ff000078-0000-0000-0000-000000000000', 'ee000020-0000-0000-0000-000000000000', 'b', 'To prevent unnecessary re-renders when a component''s props have not changed', 2),
  ('ff000079-0000-0000-0000-000000000000', 'ee000020-0000-0000-0000-000000000000', 'c', 'To split the JavaScript bundle at component boundaries and load chunks on demand', 3),
  ('ff000080-0000-0000-0000-000000000000', 'ee000020-0000-0000-0000-000000000000', 'd', 'To catch and display errors that occur anywhere inside a component subtree', 4),

  -- ee000021: Inverted pyramid
  ('ff000081-0000-0000-0000-000000000000', 'ee000021-0000-0000-0000-000000000000', 'a', 'Leading with the most important information first, then adding supporting detail below', 1),
  ('ff000082-0000-0000-0000-000000000000', 'ee000021-0000-0000-0000-000000000000', 'b', 'Starting with background context and building up to the central point at the end', 2),
  ('ff000083-0000-0000-0000-000000000000', 'ee000021-0000-0000-0000-000000000000', 'c', 'Arranging content into a visual pyramid shape on the page', 3),
  ('ff000084-0000-0000-0000-000000000000', 'ee000021-0000-0000-0000-000000000000', 'd', 'A technique for layering SEO keywords from most to least important throughout the page', 4),

  -- ee000022: Active voice
  ('ff000085-0000-0000-0000-000000000000', 'ee000022-0000-0000-0000-000000000000', 'a', 'The report was submitted by the team.', 1),
  ('ff000086-0000-0000-0000-000000000000', 'ee000022-0000-0000-0000-000000000000', 'b', 'It was decided that the feature would be postponed.', 2),
  ('ff000087-0000-0000-0000-000000000000', 'ee000022-0000-0000-0000-000000000000', 'c', 'The team submitted the report.', 3),
  ('ff000088-0000-0000-0000-000000000000', 'ee000022-0000-0000-0000-000000000000', 'd', 'The feature''s postponement was announced on Friday.', 4),

  -- ee000023: Effective headline
  ('ff000089-0000-0000-0000-000000000000', 'ee000023-0000-0000-0000-000000000000', 'a', 'It should be clever and use wordplay to reward attentive readers', 1),
  ('ff000090-0000-0000-0000-000000000000', 'ee000023-0000-0000-0000-000000000000', 'b', 'It should be as short as possible — ideally two or three words', 2),
  ('ff000091-0000-0000-0000-000000000000', 'ee000023-0000-0000-0000-000000000000', 'c', 'It should place the target SEO keyword at the very end', 3),
  ('ff000092-0000-0000-0000-000000000000', 'ee000023-0000-0000-0000-000000000000', 'd', 'It should be specific, front-loaded with key information, and understandable without the body text', 4),

  -- ee000024: F-pattern
  ('ff000093-0000-0000-0000-000000000000', 'ee000024-0000-0000-0000-000000000000', 'a', 'A five-section framework for structuring every web page', 1),
  ('ff000094-0000-0000-0000-000000000000', 'ee000024-0000-0000-0000-000000000000', 'b', 'The eye-movement pattern readers follow when scanning — across the top, then down the left side', 2),
  ('ff000095-0000-0000-0000-000000000000', 'ee000024-0000-0000-0000-000000000000', 'c', 'A technique for formatting code blocks and technical documentation', 3),
  ('ff000096-0000-0000-0000-000000000000', 'ee000024-0000-0000-0000-000000000000', 'd', 'The ideal ratio of heading text to body text on a well-structured page', 4);

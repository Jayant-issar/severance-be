ALTER TABLE "ai_reviews" DROP CONSTRAINT IF EXISTS "ai_reviews_submission_id_fkey";

ALTER TABLE "submissions" DROP CONSTRAINT IF EXISTS "submissions_assignment_id_fkey";

ALTER TABLE "submissions" DROP CONSTRAINT IF EXISTS "submissions_user_id_fkey";

ALTER TABLE "test_cases" DROP CONSTRAINT IF EXISTS "test_cases_assignment_id_fkey";

DROP TABLE IF EXISTS "ai_reviews";

DROP TABLE IF EXISTS "submissions";

DROP TABLE IF EXISTS "test_cases";

DROP TABLE IF EXISTS "assignments";

DROP TABLE IF EXISTS "users";
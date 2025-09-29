-- Reverse operations for assignment to question migration

-- 9️⃣ Remove column comment for difficulty
COMMENT ON COLUMN "questions"."difficulty" IS NULL;
-- Drop user_id foreign keys
ALTER TABLE IF EXISTS "discussions" DROP CONSTRAINT IF EXISTS discussions_user_id_fkey;
ALTER TABLE IF EXISTS "solutions" DROP CONSTRAINT IF EXISTS solutions_user_id_fkey;

-- 8️⃣ Drop new tables
DROP TABLE IF EXISTS "solutions";
DROP TABLE IF EXISTS "discussions";
DROP TABLE IF EXISTS "editorials";
DROP TABLE IF EXISTS "question_sections";

-- 4️⃣ Drop updated_at column
ALTER TABLE "questions" DROP COLUMN IF EXISTS "updated_at";

-- 2️⃣ Rename questions → assignments
ALTER TABLE "questions" RENAME TO "assignments";

-- 3️⃣ Drop new foreign keys
ALTER TABLE IF EXISTS "test_cases" DROP CONSTRAINT IF EXISTS test_cases_assignment_id_fkey;
ALTER TABLE IF EXISTS "submissions" DROP CONSTRAINT IF EXISTS submissions_assignment_id_fkey;

-- 1️⃣ Add back old foreign keys
ALTER TABLE "test_cases" ADD FOREIGN KEY ("assignment_id") REFERENCES "assignments" ("id");
ALTER TABLE "submissions" ADD FOREIGN KEY ("assignment_id") REFERENCES "assignments" ("id");
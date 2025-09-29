-- 1️⃣ Drop old foreign keys if they exist (test_cases, submissions → assignments)
ALTER TABLE IF EXISTS "test_cases" DROP CONSTRAINT IF EXISTS test_cases_assignment_id_fkey;
ALTER TABLE IF EXISTS "submissions" DROP CONSTRAINT IF EXISTS submissions_assignment_id_fkey;

-- 2️⃣ Rename assignments → questions
ALTER TABLE "assignments" RENAME TO "questions";

-- Rename columns in dependent tables
ALTER TABLE "test_cases" RENAME COLUMN "assignment_id" TO "question_id";
ALTER TABLE "submissions" RENAME COLUMN "assignment_id" TO "question_id";

-- 3️⃣ Update foreign key references to new questions table
ALTER TABLE "test_cases" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");
ALTER TABLE "submissions" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");

-- 4️⃣ Add updated_at column
ALTER TABLE "questions" ADD COLUMN "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP;

-- 5️⃣ Create flexible sections table
CREATE TABLE "question_sections" (
  "id" varchar PRIMARY KEY,
  "question_id" varchar NOT NULL REFERENCES "questions"("id") ON DELETE CASCADE,
  "title" varchar NOT NULL, -- e.g. Constraints, Examples, Explanation
  "content" text NOT NULL,
  "order" int DEFAULT 0
);

-- 6️⃣ Official Editorials
CREATE TABLE "editorials" (
  "id" varchar PRIMARY KEY,
  "question_id" varchar NOT NULL REFERENCES "questions"("id") ON DELETE CASCADE,
  "content" text NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

-- 7️⃣ Community Discussions
CREATE TABLE "discussions" (
  "id" varchar PRIMARY KEY,
  "question_id" varchar NOT NULL REFERENCES "questions"("id") ON DELETE CASCADE,
  "user_id" varchar NOT NULL, -- FK to users table
  "content" text NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE "discussions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- 8️⃣ Community Solutions
CREATE TABLE "solutions" (
  "id" varchar PRIMARY KEY,
  "question_id" varchar NOT NULL REFERENCES "questions"("id") ON DELETE CASCADE,
  "user_id" varchar NOT NULL, -- FK to users table
  "content" text NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE "solutions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- 9️⃣ Column comment for difficulty
COMMENT ON COLUMN "questions"."difficulty" IS 'easy | medium | hard';

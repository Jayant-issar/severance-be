CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "assignments" (
  "id" varchar PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" text NOT NULL,
  "difficulty" varchar NOT NULL,
  "tags" text,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "test_cases" (
  "id" varchar PRIMARY KEY,
  "assignment_id" varchar NOT NULL,
  "input" text NOT NULL,
  "expected_output" text NOT NULL,
  "is_hidden" boolean DEFAULT true,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "submissions" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "assignment_id" varchar NOT NULL,
  "code" text NOT NULL,
  "language" varchar NOT NULL,
  "status" varchar DEFAULT 'pending',
  "runtime_ms" int,
  "memory_kb" int,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "ai_reviews" (
  "id" varchar PRIMARY KEY,
  "submission_id" varchar NOT NULL,
  "feedback" text NOT NULL,
  "score" int,
  "review_agent" varchar,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

COMMENT ON COLUMN "assignments"."difficulty" IS 'easy | medium | hard';

COMMENT ON COLUMN "submissions"."status" IS 'pending | running | passed | failed';

COMMENT ON COLUMN "ai_reviews"."score" IS '0-100 optional';

ALTER TABLE "test_cases" ADD FOREIGN KEY ("assignment_id") REFERENCES "assignments" ("id");

ALTER TABLE "submissions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "submissions" ADD FOREIGN KEY ("assignment_id") REFERENCES "assignments" ("id");

ALTER TABLE "ai_reviews" ADD FOREIGN KEY ("submission_id") REFERENCES "submissions" ("id");
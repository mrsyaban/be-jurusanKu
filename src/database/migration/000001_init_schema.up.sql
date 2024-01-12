CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "nick" varchar NOT NULL DEFAULT '',
  "role" varchar NOT NULL CHECK (role IN ('mentor', 'student')),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "last_test_date" timestamp DEFAULT NULL
);

CREATE TABLE "majors" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "desc" varchar NOT NULL,
  "image_url" varchar NOT NULL,
  "interest_num" integer NOT NULL DEFAULT 0
);

CREATE TABLE "courses" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "desc" varchar NOT NULL,
  "major_id" bigserial NOT NULL,
  "image_url" varchar NOT NULL,
  "price" integer NOT NULL,
  "syllabus" text NOT NULL
);

CREATE TABLE "projects" (
  "id" bigserial PRIMARY KEY,
  "course_id" bigserial NOT NULL,
  "image_url" varchar NOT NULL,
  "content_url" varchar NOT NULL
);

CREATE TABLE "enrollment" (
  "user_email" varchar NOT NULL,
  "course_id" bigserial  NOT NULL,
  PRIMARY KEY ("user_email", "course_id")
);

CREATE TABLE "aptitude_question" (
  "id" bigserial PRIMARY KEY,
  "question" varchar NOT NULL
);

CREATE TABLE "aptitude_test" (
  "question_id" bigserial,
  "user_id" bigserial,
  "answer" varchar NOT NULL,
  PRIMARY KEY ("question_id", "user_id")
);

CREATE TABLE "material" (
  "id" bigserial PRIMARY KEY,
  "type" varchar CHECK (type IN ('video', 'document')),
  "title" varchar NOT NULL,
  "content_url" varchar NOT NULL,
  "course_id" bigserial NOT NULL
);

CREATE TABLE "recommendation" (
  "user_id" bigserial NOT NULL,
  "major_id" bigserial NOT NULL,
  "percentage" integer NOT NULL,
  PRIMARY KEY ("user_id", "major_id")
);

CREATE TABLE "courses_majors" (
  "course_id" bigserial NOT NULL,
  "major_id" bigserial NOT NULL,
  PRIMARY KEY ("course_id", "major_id")
);

ALTER TABLE "material" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "projects" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "enrollment" ADD FOREIGN KEY ("user_email") REFERENCES "users" ("email");

ALTER TABLE "enrollment" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "aptitude_test" ADD FOREIGN KEY ("question_id") REFERENCES "aptitude_question" ("id");

ALTER TABLE "aptitude_test" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "recommendation" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "recommendation" ADD FOREIGN KEY ("major_id") REFERENCES "majors" ("id");

ALTER TABLE "courses_majors" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "courses_majors" ADD FOREIGN KEY ("major_id") REFERENCES "majors" ("id");

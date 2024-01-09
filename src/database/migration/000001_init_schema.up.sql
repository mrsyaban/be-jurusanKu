CREATE TABLE users (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE TABLE majors (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "desc" varchar NOT NULL
);

CREATE TABLE courses (
  "id" bigserial PRIMARY KEY,
  "major_id" integer NOT NULL,
  "desc" varchar NOT NULL
);

ALTER TABLE "courses" ADD FOREIGN KEY ("major_id") REFERENCES "majors" ("id");

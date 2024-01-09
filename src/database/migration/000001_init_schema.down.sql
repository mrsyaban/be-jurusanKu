ALTER TABLE "courses" DROP CONSTRAINT "courses_major_id_fkey";

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS majors;
DROP TABLE IF EXISTS courses;
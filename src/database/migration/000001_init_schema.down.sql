ALTER TABLE "material" DROP CONSTRAINT "material_course_id_fkey";

ALTER TABLE "projects" DROP CONSTRAINT "projects_course_id_fkey";

ALTER TABLE "enrollment" DROP CONSTRAINT "enrollment_user_id_fkey";

ALTER TABLE "enrollment" DROP CONSTRAINT "enrollment_course_id_fkey";

ALTER TABLE "aptitude_test" DROP CONSTRAINT "aptitude_test_question_id_fkey";

ALTER TABLE "aptitude_test" DROP CONSTRAINT "aptitude_test_user_id_fkey";

ALTER TABLE "recommendation" DROP CONSTRAINT "recommendation_user_id_fkey";

ALTER TABLE "recommendation" DROP CONSTRAINT "recommendation_major_id_fkey";

ALTER TABLE "courses_majors" DROP CONSTRAINT "courses_majors_course_id_fkey";

ALTER TABLE "courses_majors" DROP CONSTRAINT "courses_majors_major_id_fkey";

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS majors;
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS enrollment;
DROP TABLE IF EXISTS aptitude_question;
DROP TABLE IF EXISTS aptitude_test;
DROP TABLE IF EXISTS material;
DROP TABLE IF EXISTS recommendation;
DROP TABLE IF EXISTS courses_majors;
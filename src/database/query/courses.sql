-- name: CreateCourse :one
INSERT INTO courses (
    title,
    "desc",
    "major_id"
) VALUES (
    'Computer Science',
    'The study of computers and their applications',
    1
) RETURNING *;

-- name: GetCourses :many
SELECT * FROM courses;

-- name: GetCourseByMajorId :one
SELECT * FROM courses WHERE major_id = @major_id;
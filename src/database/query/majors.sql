-- name: CreateMajor :one
INSERT INTO majors (
    title,
    "desc"
) VALUES (
    'Computer Science',
    'The study of computers and their applications'
) RETURNING *;

-- name: GetMajors :many
SELECT * FROM majors;
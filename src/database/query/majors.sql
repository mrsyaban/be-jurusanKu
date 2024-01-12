-- name: CreateDummyMajor :many
INSERT INTO majors (
    title,
    "desc",
    "interest_num"
) VALUES (
    'Teknik Informatika',
    'The study of computers and their applications',
    1100
) RETURNING *;


-- name: GetMajors :many
SELECT * FROM majors;
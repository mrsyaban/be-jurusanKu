-- name: CreateDummyMajor :many
INSERT INTO majors (
    title,
    "desc",
    "interest_num",
    image_url
) VALUES (
    'Teknik Informatika',
    'The study of computers and their applications',
    1100,
    'https://sample.dummy'
) RETURNING *;


-- name: GetMajors :many
SELECT * FROM majors;
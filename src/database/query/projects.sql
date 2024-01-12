-- name: GetProjectByCourseId :many
SELECT * 
FROM projects 
WHERE course_id = @course_id;

-- name: CreateDummyProject :one
INSERT INTO projects (
    course_id,
    image_url,
    content_url
) VALUES (
    1,
    'https://sample.dummy',
    'https://sample.dummy'
) RETURNING *;
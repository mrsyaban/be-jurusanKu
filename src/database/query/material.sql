-- name: GetMaterialByCourseId :one
SELECT * 
FROM material 
WHERE course_id = @course_id;

-- name: CreateDummyMaterial :one
INSERT INTO material (
    type,
    title,
    content_url,
    course_id
) VALUES (
    'video',
    'Introduction to Computer Science',
    'https://sample.dummy',
    1
) RETURNING *;
-- name: GetProjectByCourseId :many
SELECT * 
FROM projects 
WHERE course_id = @course_id;
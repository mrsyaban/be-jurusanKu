-- name: GetMaterialByCourseId :one
SELECT * 
FROM material 
WHERE course_id = @course_id;
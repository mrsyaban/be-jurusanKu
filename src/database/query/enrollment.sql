-- name: AddEnrollment :one
INSERT INTO enrollment (
    user_email,
    course_id
) VALUES (
    @user_email,
    @course_id
) RETURNING *;

-- name: GetActiveCourse :many
SELECT course_id 
FROM enrollment 
WHERE user_email = @user_email;
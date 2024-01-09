-- name: createUser :one
INSERT INTO users (
    username,
    password,
    role
) VALUES (
    'admin',
    'password',
    'admin'
) RETURNING *;
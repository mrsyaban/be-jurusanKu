-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    role
) VALUES (
    'admin',
    'password',
    'admin'
) RETURNING *;
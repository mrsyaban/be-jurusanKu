-- name: CreateDummyUser :one
INSERT INTO users (
    username,
    password,
    role
) VALUES (
    'adminDummy',
    'password',
    'admin'
) RETURNING *;

-- name: RegisterUser :one
INSERT INTO users (
    username,
    password,
    role
) VALUES (
    @username,
    @password,
    @role
) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = @username;
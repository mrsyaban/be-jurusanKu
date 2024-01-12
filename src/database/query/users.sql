-- name: CreateDummyUser :one
INSERT INTO users (
    email,
    password,
    role,
    name,
    nick
) VALUES (
    'jurusanku@gmail.com',
    'password',
    'student',
    'jurusanku',
    ''
) RETURNING *;

-- name: RegisterUser :one
INSERT INTO users (
    email,
    password,
    role,
    name
) VALUES (
    @email,
    @password,
    @role,
    @name
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * 
FROM users 
WHERE email = @email;

-- name: UpdateProfile :one
UPDATE users
SET
    name = @name,
    "nick" = @nick
WHERE email = @email
RETURNING *;
-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  password,
  role,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
set name = $2,
email = $3,
password = $4,
role = $5,
updated_at = $6
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

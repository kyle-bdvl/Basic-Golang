-- name: CreateUser :one
INSERT INTO user_table ( username, created_date)
VALUES ( $1, NOW())
RETURNING id, username, created_date; 

-- name: GetUser :one
SELECT *
FROM user_table
WHERE id = $1
LIMIT 1;

-- name: GetAllUsers :many
SELECT * 
FROM user_table;
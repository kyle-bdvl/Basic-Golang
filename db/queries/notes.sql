-- name: ListNotesByUser :many
SELECT *
FROM note_table
WHERE userid = $1;

-- name: CreateNote :one 
INSERT INTO note_table( created_date, description_of_note, userid )
VALUES (NOW(),$1,$2)
RETURNING *;

-- name: UpdateTitle :exec
UPDATE note_table
SET title = $2
WHERE id = $1;
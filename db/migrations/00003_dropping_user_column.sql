-- +goose Up
ALTER TABLE note_table
DROP COLUMN IF EXISTS username;

-- +goose Down
ALTER TABLE note_table
ADD COLUMN username varchar(256);
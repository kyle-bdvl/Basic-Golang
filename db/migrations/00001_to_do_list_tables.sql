-- +goose Up
-- SELECT 'up SQL query';
CREATE TABLE User_table (
    id uuid  PRIMARY KEY,
    username varchar(256) unique,
    created_date timestamp
);
CREATE TABLE Note_table (
    id uuid unique PRIMARY KEY,
    userid Uuid REFERENCES User_table(id), 
    username varchar(256), 
    created_date timestamp, 
    title varchar(256),
    description_of_note varchar(256)
);

-- +goose Down
-- SELECT 'down SQL query';
DROP TABLE IF EXISTS User_table;
DROP TABLE IF EXISTS Note_table;

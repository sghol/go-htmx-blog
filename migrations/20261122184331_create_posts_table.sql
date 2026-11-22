-- migrations/20261122184331_create_posts_table.sql

-- +goose UP
CREATE TABLE posts (
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    title        TEXT NOT NULL,
    body         TEXT NOT NULL,
    is_published BOOLEAN NOT NULL DEFAULT 0,
    created_at   DATETIME NOT NULL,
    updated_at   DATETIME NOT NULL
);


-- +goose DOWN
DROP TABLE posts;

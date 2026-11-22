-- queries/posts.sql

-- name: ListPosts :many
SELECT * FROM posts ORDER BY created_at DESC;


-- name: CreatePost :one
INSERT INTO posts (title, body, is_published, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)  
RETURNING *;

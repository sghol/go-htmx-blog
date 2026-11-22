-- queries/posts.sql

-- name: ListPosts :many
SELECT * FROM posts ORDER BY created_at DESC;

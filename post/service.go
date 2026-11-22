package post

import (
	"context"
	"database/sql"
	"time"

	"github.com/sghol/go-htmx-blog/db"
)

type Service struct {
	db *sql.DB
	q  *db.Queries
}

func NewService(sqlDB *sql.DB) *Service {
	return &Service{db: sqlDB, q: db.New(sqlDB)}
}

// ListPosts
func (s *Service) ListPosts(ctx context.Context) ([]db.Post, error) {
	return s.q.ListPosts(ctx)
}

type CreatePostInput struct {
	Title       string
	Body        string
	IsPublished bool
}

// CreatePost create a new post
func (s *Service) CreatePost(ctx context.Context, inp CreatePostInput) (db.Post, error) {
	return s.q.CreatePost(ctx, db.CreatePostParams{
		Title:       inp.Title,
		Body:        inp.Body,
		IsPublished: inp.IsPublished,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
}

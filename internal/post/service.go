package post

import (
	"context"
	"database/sql"
	"github.com/sghol/go-htmx-blog/db"
)

type Service struct {
	db *sql.DB
	q *db.Queries
}

func NewService(sqlDB *sql.DB) *Service {
	return &Service{db: sqlDB, q: db.New(sqlDB)}
}

// ListPosts 
func (s *Service) ListPosts(ctx context.Context) ([]db.Post, error){
	return s.q.ListPosts(ctx)

}

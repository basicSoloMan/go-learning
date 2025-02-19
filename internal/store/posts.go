package store

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
)

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID, pq.Array(post.Tags)).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	return err
}

func (s *PostStore) GetById(ctx context.Context, postId int64) (*Post, error) {
	query := `
		SELECT id, title, user_id, content, created_at, tags, updated_at
		FROM posts
		WHERE id = $1
	`

	var post Post
	err := s.db.QueryRowContext(ctx, query, postId).Scan(&post.ID, &post.Title, &post.UserID, &post.Content, &post.CreatedAt, pq.Array(&post.Tags), &post.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

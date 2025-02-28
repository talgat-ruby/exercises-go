package post

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

func (m *Post) PostsCreate(ctx context.Context, model *Model) (*Model, error) {
	log := m.logger.With("post_method", "PostsCreate")

	stmt := `INSERT INTO posts (title, content, category, tags) VALUES ($1, $2, $3, $4) RETURNING id, title, content, category, tags, created_at, updated_at`

	row := m.db.QueryRowContext(ctx, stmt, model.Title, model.Content, model.Category, pq.Array(model.Tags))
	if row.Err() != nil {
		return nil, row.Err()
	}

	result := Model{}

	if err := row.Scan(&result.ID, &result.Title, &result.Content, &result.Category, &result.Tags, &result.CreatedAt, &result.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.ErrorContext(ctx, "failed create to db", "error", err)
		return nil, err
	}

	return &result, nil
}

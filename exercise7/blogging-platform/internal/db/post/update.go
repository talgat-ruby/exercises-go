package post

import (
	"context"

	"github.com/lib/pq"
)

func (m *Post) PostsUpdate(ctx context.Context, id int64, model *Model) (*Model, error) {
	log := m.logger.With("post_method", "PostsUpdate")

	stmt := `UPDATE posts SET title = $2, content = $3, category = $4, tags = $5 WHERE id = $1 RETURNING id, title, content, category, tags, created_at, updated_at`

	row := m.db.QueryRowContext(ctx, stmt, id, model.Title, model.Content, model.Category, pq.Array(model.Tags))
	if row.Err() != nil {
		return nil, row.Err()
	}

	result := Model{}

	if err := row.Scan(&result.ID, &result.Title, &result.Content, &result.Category, &result.Tags, &result.CreatedAt, &result.UpdatedAt); err != nil {
		log.ErrorContext(ctx, "failed update", "error", err)
		return nil, err
	}

	return &result, nil
}

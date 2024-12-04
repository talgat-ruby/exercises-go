package post

import (
	"context"
)

func (m *Post) FindPost(ctx context.Context, id int64) (*ModelPost, error) {
	log := m.logger.With("method", "FindPost")

	stmt := `
SELECT id, title, content, category, tags, created_at, updated_at 
FROM post
WHERE id = $1
`

	row := m.db.QueryRowContext(ctx, stmt, id)

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to query table post", "error", err)
		return nil, err
	}

	post := ModelPost{}

	if err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Category,
		&post.Tags,
		&post.CreatedAt,
		&post.UpdatedAt,
	); err != nil {
		log.ErrorContext(ctx, "fail to scan post", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success query table post")
	return &post, nil
}

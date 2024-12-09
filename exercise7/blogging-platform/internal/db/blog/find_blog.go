package blog

import (
	"context"
	"database/sql"
	"fmt"
)

func (m *Blog) FindBlogPost(ctx context.Context, id int64) (*BlogPost, error) {
	log := m.logger.With("method", "FindBlogPost")

	stmt := `
ELECT id, title, content, category, tags, created_at, updated_at 
FROM blog_posts
WHERE id = $1
`

	row := m.db.QueryRowContext(ctx, stmt, id)

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "failed to query blog_posts table", "error", err)
		return nil, fmt.Errorf("querying blogs post: %w", err)
	}

	blog := BlogPost{}

	if err := row.Scan(
		&blog.ID,
		&blog.Title,
		&blog.Content,
		&blog.Category,
		&blog.CreatedAt,
		&blog.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			log.InfoContext(ctx, "blogs post not found", "id", id)
			return nil, nil // or return specific error like ErrBlogNotFound
		}
		log.ErrorContext(ctx, "failed to scan blogs post", "error", err)
		return nil, fmt.Errorf("scanning blogs post: %w", err)
	}

	log.InfoContext(ctx, "success query table movie", "id", id)
	return &blog, nil
}

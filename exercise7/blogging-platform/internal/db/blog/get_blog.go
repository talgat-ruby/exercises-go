package blog

import (
	"context"
)

func (b *Blog) GetBlog(ctx context.Context, id int64) (*ModelBlogs, error) {
	log := b.logger.With("method", "GetBlog")

	stmt := `
SELECT id, title, description, created_at, updated_at 
FROM blogs
WHERE id = $1
`

	row := b.db.QueryRowContext(ctx, stmt, id)

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to query table blogs", "error", err)
		return nil, err
	}

	blog := ModelBlogs{}

	if err := row.Scan(
		&blog.ID,
		&blog.Title,
		&blog.Description,
		&blog.CreatedAt,
		&blog.UpdatedAt,
	); err != nil {
		log.ErrorContext(ctx, "fail to scan blogs", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success query table blogs")
	return &blog, nil
}

package blog

import (
	"context"
	"database/sql"
	"errors"
)

func (b *Blog) CreateBlog(ctx context.Context, insertData *ModelBlogs) (*ModelBlogs, error) {
	log := b.logger.With("method", "CreateBlog")

	stmt := `
INSERT INTO blogs (title, description)
VALUES ($1, $2)
RETURNING id, title, description, created_at, updated_at
`

	row := b.db.QueryRowContext(ctx, stmt, insertData.Title, insertData.Description)

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to insert to table blogs", "error", err)
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
		if errors.Is(err, sql.ErrNoRows) {
			log.ErrorContext(ctx, "no values was found", "error", err)
			return nil, nil
		}
		log.ErrorContext(ctx, "fail to scan blog", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success insert to table blogs")
	return &blog, nil
}

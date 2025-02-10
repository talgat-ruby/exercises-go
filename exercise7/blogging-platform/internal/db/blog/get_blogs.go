package blog

import "context"

func (b *Blog) GetBlogs(ctx context.Context) ([]ModelBlogs, error) {
	log := b.logger.With("db", "GetBlogs")

	blogs := make([]ModelBlogs, 0)

	stmt := `
SELECT id, title, description, created_at, updated_at 
FROM blogs
`

	rows, err := b.db.QueryContext(ctx, stmt)
	if err != nil {
		log.ErrorContext(ctx, "fail to query table blogs", "error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		blog := ModelBlogs{}

		if err := rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Description,
			&blog.CreatedAt,
			&blog.UpdatedAt,
		); err != nil {
			log.ErrorContext(ctx, "fail to scan blog", "error", err)
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	if err := rows.Err(); err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success query table blogs", "blogs")
	return blogs, nil
}

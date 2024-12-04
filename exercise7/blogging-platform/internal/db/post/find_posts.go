package post

import (
	"context"
)

func (m *Post) FindPosts(ctx context.Context) ([]ModelPost, error) {
	log := m.logger.With("method", "FindPosts")

	posts := make([]ModelPost, 0)

	stmt := `
SELECT id, title, content, category, tags, created_at, updated_at 
FROM post
`

	rows, err := m.db.QueryContext(ctx, stmt)
	if err != nil {
		log.ErrorContext(ctx, "fail to query table post", "error", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		post := ModelPost{}

		if err := rows.Scan(
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

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success query table post")
	return posts, nil
}

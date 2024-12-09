package blog

import (
	"context"
	"encoding/json"
	"fmt"
)

func (m *Blog) FindBlogPosts(ctx context.Context, offset, limit int) ([]BlogPost, error) {
	log := m.logger.With("method", "FindBlogPosts")

	posts := make([]BlogPost, 0)

	stmt := `
    SELECT 
        b.id,
        b.title,
        b.content,
        b.category,
        b.created_at,
        b.updated_at,
        COALESCE(json_agg(json_build_object(
            'id', t.id,
            'name', t.name
        )) FILTER (WHERE t.id IS NOT NULL), '[]'::json) as tags
    FROM blog_posts b
    LEFT JOIN blog_posts_tags bt ON b.id = bt.blog_post_id
    LEFT JOIN tags t ON bt.tag_id = t.id
    GROUP BY b.id, b.title, b.content, b.category, b.created_at, b.updated_at
    ORDER BY b.created_at DESC
    OFFSET $1 LIMIT $2;
    `

	rows, err := m.db.QueryContext(ctx, stmt, offset, limit)
	if err != nil {
		log.ErrorContext(ctx, "failed to query blog posts", "error", err)
		return nil, fmt.Errorf("querying blog posts: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		blog := BlogPost{}
		var tagsJSON []byte

		if err := rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Content,
			&blog.Category,
			&blog.CreatedAt,
			&blog.UpdatedAt,
			&tagsJSON,
		); err != nil {
			log.ErrorContext(ctx, "failed to scan blog post", "error", err)
			return nil, fmt.Errorf("scanning blog post: %w", err)
		}

		// Initialize empty tags slice
		blog.Tags = []Tag{}

		// Unmarshal tags if we have any
		if err := json.Unmarshal(tagsJSON, &blog.Tags); err != nil {
			log.ErrorContext(ctx, "failed to unmarshal tags", "error", err)
			return nil, fmt.Errorf("unmarshaling tags: %w", err)
		}

		posts = append(posts, blog)
	}

	if err := rows.Err(); err != nil {
		log.ErrorContext(ctx, "failed to iterate rows", "error", err)
		return nil, fmt.Errorf("iterating rows: %w", err)
	}

	log.InfoContext(ctx, "successfully queried blog posts",
		"count", len(posts),
		"offset", offset,
		"limit", limit,
	)

	return posts, nil
}

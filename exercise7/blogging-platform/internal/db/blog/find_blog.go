package blog

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

func (m *Blog) FindBlogPost(ctx context.Context, id int64) (*BlogPost, error) {
	log := m.logger.With("method", "FindBlogPost")

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
WHERE b.id = $1
GROUP BY b.id, b.title, b.content, b.category, b.created_at, b.updated_at;
`

	var tagsJSON []byte
	blog := BlogPost{}

	if err := m.db.QueryRowContext(ctx, stmt, id).Scan(
		&blog.ID,
		&blog.Title,
		&blog.Content,
		&blog.Category,
		&blog.CreatedAt,
		&blog.UpdatedAt,
		&tagsJSON,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrBlogNotFound
		}
		log.ErrorContext(ctx, "failed to scan blog post",
			"error", err,
			"id", id,
		)
		return nil, fmt.Errorf("scanning blog post: %w", err)
	}

	// Initialize empty tags slice
	blog.Tags = []Tag{}

	// Unmarshal the JSON tags data
	if err := json.Unmarshal(tagsJSON, &blog.Tags); err != nil {
		log.ErrorContext(ctx, "failed to unmarshal tags",
			"error", err,
			"id", id,
		)
		return nil, fmt.Errorf("unmarshaling tags: %w", err)
	}

	log.InfoContext(ctx, "success query table blog_post", "id", id)
	return &blog, nil
}

var (
	ErrBlogNotFound = errors.New("blog post not found")
	ErrInvalidID    = errors.New("invalid blog post ID")
)

/*

type BlogPost struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Category  string    `json:"category" db:"category"`
	Tags      []string  `json:"tags" db:"tags"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
*/

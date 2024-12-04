package post

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

func (m *Post) CreatePost(ctx context.Context, insertData *ModelPost) (*ModelPost, error) {
	log := m.logger.With("method", "CreatePost")

	stmt := `
INSERT INTO post (title, content, category, tags)
VALUES ($1, $2, $3, $4)
RETURNING id, title, content, category, tags, created_at, updated_at
`

	row := m.db.QueryRowContext(ctx, stmt, insertData.Title, insertData.Content, insertData.Category, pq.Array(insertData.Tags))

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to insert to table post", "error", err)
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
		if errors.Is(err, sql.ErrNoRows) {
			log.ErrorContext(ctx, "no values was found", "error", err)
			return nil, nil
		}
		log.ErrorContext(ctx, "fail to scan post", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success insert to table post")
	return &post, nil
}

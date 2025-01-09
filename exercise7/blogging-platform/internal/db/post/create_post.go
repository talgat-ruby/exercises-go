package post

import (
	"context"
	"database/sql"
	"errors"
)

func (m *Post) CreatePost(ctx context.Context, insertData *ModelPost) (*ModelPost, error) {
	log := m.logger.With("method", "CreatePost")

	stmt := `
INSERT INTO post (title, description, posterUrl)
VALUES ($1, $2, $3)
RETURNING id, title, description, posterUrl, created_at, updated_at
`

	row := m.db.QueryRowContext(ctx, stmt, insertData.Title, insertData.Description, insertData.PosterURL)

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to insert to table post", "error", err)
		return nil, err
	}

	post := ModelPost{}

	if err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Description,
		&post.PosterURL,
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

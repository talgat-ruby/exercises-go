package post

import (
	"context"
)

func (m *Post) PostsDelete(ctx context.Context, id int64) error {
	log := m.logger.With("post_method", "PostsDelete")

	stmt := `DELETE FROM posts WHERE id = $1 RETURNING id`

	row := m.db.QueryRowContext(ctx, stmt, id)
	if row.Err() != nil {
		log.ErrorContext(ctx, "failed delete db", "error", row.Err())
		return row.Err()
	}

	err := row.Scan(&id)

	if err != nil {
		return err
	}

	return nil
}

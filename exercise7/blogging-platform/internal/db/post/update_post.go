package post

import (
	"context"
	"fmt"
)

func (m *Post) UpdatePost(ctx context.Context, id int64, insertData *ModelPost) error {
	log := m.logger.With("method", "UpdatePost", "id", id)

	stmt := `
UPDATE post
SET title = $2, description = $3, posterUrl = $4
WHERE id = $1
`

	res, err := m.db.ExecContext(ctx, stmt, id, insertData.Title, insertData.Description, insertData.PosterURL)
	if err != nil {
		log.ErrorContext(ctx, "fail to update the table post", "error", err)
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "fail to update from the table post", "error", err)
		return err
	}

	if num == 0 {
		log.WarnContext(ctx, "post with id was not found", "id", id)
		return fmt.Errorf("post with id was not found")
	}

	log.InfoContext(ctx, "success update the table post")
	return nil
}

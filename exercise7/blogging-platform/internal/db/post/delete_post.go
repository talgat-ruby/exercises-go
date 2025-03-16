package post

import (
	"context"
	"fmt"
)

func (m *Post) DeletePost(ctx context.Context, id int64) error {
	log := m.logger.With("method", "DeletePost", "id", id)

	stmt := `
DELETE FROM post
WHERE id = $1
`

	res, err := m.db.ExecContext(ctx, stmt, id)
	if err != nil {
		log.ErrorContext(ctx, "fail to delete from the table post", "error", err)
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "fail to delete from the table post", "error", err)
		return err
	}

	if num == 0 {
		log.WarnContext(ctx, "post with id was not found", "id", id)
		return fmt.Errorf("post with id was not found")
	}

	log.InfoContext(ctx, "success delete from the table post")
	return nil
}

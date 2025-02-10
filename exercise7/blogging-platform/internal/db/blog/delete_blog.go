package blog

import (
	"context"
	"fmt"
)

func (b *Blog) DeleteBlog(ctx context.Context, id int64) error {
	log := b.logger.With("method", "DeleteBlog", "id", id)

	stmt := `
DELETE FROM blogs
WHERE id = $1
`

	res, err := b.db.ExecContext(ctx, stmt, id)
	if err != nil {
		log.ErrorContext(ctx, "fail to delete from the table blogs", "error", err)
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "fail to delete from the table blogs", "error", err)
		return err
	}

	if num == 0 {
		log.WarnContext(ctx, "blog with id was not found", "id", id)
		return fmt.Errorf("blog with id was not found")
	}

	log.InfoContext(ctx, "success delete from the table blogs")
	return nil
}

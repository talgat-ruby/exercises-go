package blog

import (
	"context"
	"fmt"
)

func (b *Blog) UpdateBlog(ctx context.Context, id int64, insertData *ModelBlogs) error {
	log := b.logger.With("method", "UpdateBlog", "id", id)

	stmt := `
UPDATE blogs
SET title = $2, description = $3
WHERE id = $1
`

	res, err := b.db.ExecContext(ctx, stmt, id, insertData.Title, insertData.Description)
	if err != nil {
		log.ErrorContext(ctx, "fail to update the table blogs", "error", err)
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "fail to update from the table blogs", "error", err)
		return err
	}

	if num == 0 {
		log.WarnContext(ctx, "blog with id was not found", "id", id)
		return fmt.Errorf("blog with id was not found")
	}

	log.InfoContext(ctx, "success update the table blogs")
	return nil
}

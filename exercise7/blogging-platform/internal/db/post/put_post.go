package post

import (
	"context"
	"fmt"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

func (p Post) UpdatePost(ctx context.Context, id int64, body models.Blog) error {
	log := p.logger.With("method", "updatePost", "id", id)
	stmt := `
UPDATE posts
SET title = $2, content = $3, category = $4, tags = $5
WHERE id = $1
	`
	res, err := p.db.ExecContext(ctx, stmt, id, body.Title, body.Content, body.Category, body.Tags)
	if err != nil {
		log.ErrorContext(ctx, "fail updated the table posts", "error", err)
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "fail updated the table posts", "error", err)
		return err
	}
	if num == 0 {
		log.WarnContext(ctx, "post with id was not found", "id", id)
		return fmt.Errorf("post with id was not found")
	}
	log.InfoContext(ctx, "succes update")
	return nil
}

package post

import "context"

func (p Post) DeletePost(id int64, ctx context.Context) error {
	log := p.logger.With("metgod", "deletePost", "id", id)
	stmt := `
	DELETE FROM posts
	WHERE id = $1
	`
	res, err := p.db.ExecContext(ctx, stmt, id)
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
		log.WarnContext(ctx, "movie with id was not found", "error", err)
		return err
	}
	log.InfoContext(ctx, "succes delete")
	return nil
}

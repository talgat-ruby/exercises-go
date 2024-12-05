package post

import (
	"context"
	"fmt"
)

func (p *Post) DeletePost(ctx context.Context, id int) error {
	log := p.logger.With("method", "DeletePost")

	queryDeletePost := `DELETE FROM post WHERE id = $1`
	deletePost, err := p.db.ExecContext(ctx, queryDeletePost, id)
	if err != nil {
		log.ErrorContext(ctx, "failed to delete post", "error", err)
		return err
	}

	countPost, err := deletePost.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "failed to get rows affected", "error", err)
		return err
	}

	if countPost == 0 {
		log.ErrorContext(ctx, "post not found", "post_id", id)
		return fmt.Errorf("post with id %d not found", id)
	}

	log.InfoContext(ctx, "success delete post")

	return nil
}

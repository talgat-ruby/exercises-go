package post

import (
	"context"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

func (p *Post) ServiceGetNumberPost(ctx context.Context, id int) (error, *models.Blog) {
	log := p.logger.With("method", "GetPost")
	stmt := `
	SELECT id, title, description, posterUrl, created_at, updated_at 
	FROM movie
	WHERE id = $1
	`
	row := p.db.QueryRowContext(ctx, stmt, id)
	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to query table post", "error", err)
		return err, nil
	}
	post := models.Blog{}
	if err := row.Scan(
		&post.Id,
		&post.Title,
		&post.Content,
		&post.Category,
		&post.Tags,
		&post.CreatedAt,
		&post.UpdatedAt,
	); err != nil {
		log.ErrorContext(ctx, "fail to scan post", "error", err)
		return err, nil
	}

	log.InfoContext(ctx, "succes query table posts")
	return nil, &post
}

package post

import (
	"context"
	"database/sql"
	"errors"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

func (p *Post) NewPostCreator(ctx context.Context, body models.Blog) (*models.Blog, error) {
	log := p.logger.With("method", "CreatePost")
	res := `
	INSERT INTO posts (
	title,
	content,
	category,
	tags
	
	)
	VALUES ($1, $2, $3, $4)
	`
	row := p.db.QueryRowContext(ctx, res, body.Title, body.Content, body.Category, body.Tags)
	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to insert to table post", err)
		return nil, err
	}
	post := models.Blog{}
	if err := row.Scan(
		&post.Id,
		&post.Title,
		&post.Content,
		&post.Category,
		post.Tags,
		post.UpdatedAt,
		post.CreatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.ErrorContext(ctx, "no values was found", "erros", err)
			return nil, nil
		}
		log.ErrorContext(ctx, "fail to scan post", "error", err)
		return nil, err
	}
	// _, err = db.Exec(res, body.Title, body.Content, body.Category, body.Tags)
	// return err

	log.InfoContext(ctx, "succes insert")
	return &post, nil
}

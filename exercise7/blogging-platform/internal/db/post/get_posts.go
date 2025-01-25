package post

import (
	"context"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

func (p Post) DBGetPosts(ctx context.Context) ([]models.Blog, error) {
	log := p.logger.With("method", "GetPost")
	stmt := `
	SELECT id, title, content, category, tags, created_at, updated_at 
	FROM posts
	`

	rows, err := p.db.QueryContext(ctx, stmt)
	if err != nil {
		log.ErrorContext(ctx, "fail to query table post", "error", err)
		return nil, err
	}

	posts := []models.Blog{}
	post := models.Blog{}
	for rows.Next() {
		if err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.Category,
			&post.Tags,
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil {
			log.ErrorContext(ctx, "fail to scan post", "error", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	log.InfoContext(ctx, "succes query table posts")
	return posts, nil
}

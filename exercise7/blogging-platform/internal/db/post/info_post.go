package post

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
)

func (p *Post) GetInformationOfPost(ctx context.Context, id int) (blog.Post, error) {
	log := p.logger.With("method", "GetInformationOfPost")

	query := `
		SELECT p.id, p.title, p.content, p.created_at, p.updated_at, 
		       c.id, c.name,  
		       t.id, t.name
		FROM post p
		JOIN category c ON p.id_category = c.id
		LEFT JOIN post_tags pt ON pt.id_post = p.id
		LEFT JOIN tag t ON pt.id_tag = t.id
		WHERE p.id = $1
	`

	rows, err := p.db.QueryContext(ctx, query, id)
	if err != nil {
		log.ErrorContext(ctx, "fail to get post information", "error", err)
		return blog.Post{}, err
	}
	defer rows.Close()

	var post blog.Post
	var category blog.Category
	var tags []*blog.Tag
	postFound := false
	for rows.Next() {
		var tagID sql.NullInt64
		var tagName sql.NullString

		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt,
			&category.ID, &category.Name, &tagID, &tagName)
		if err != nil {
			log.ErrorContext(ctx, "failed to scan row", "error", err)
			return blog.Post{}, err
		}

		post.Category = &category

		if tagID.Valid && tagName.Valid {
			tags = append(tags, &blog.Tag{ID: int(tagID.Int64), Name: tagName.String})
		}

		postFound = true
	}

	if !postFound {
		log.ErrorContext(ctx, "post not found", "post_id", id)
		return blog.Post{}, fmt.Errorf("post with id %d not found", id)
	}

	post.Tags = tags

	if len(tags) == 0 {
		log.InfoContext(ctx, "no tags found for the post", "post_id", id)
	}

	log.InfoContext(ctx, "success get information of post")
	return post, nil
}

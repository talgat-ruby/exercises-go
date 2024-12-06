package post

import (
	"context"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

func (p Post) DBGetPosts(ctx context.Context) ([]models.Blog, error) {
	// db, err := New()
	// if err != nil {
	// 	return nil, err
	// }
	var res []models.Blog
	// res, err := db.Exec("SELECT  FROM posts")
	rows, err := p.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post models.Blog
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Category, &post.Tags, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

package category

import (
	"context"
	"database/sql"
	"sort"
	"strconv"
	"strings"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
)

func (c *Category) GetAllPostsOfCategory(ctx context.Context, id int) ([]blog.Post, error) {
	log := c.logger.With("method", "GetAllPostsOfCategory")

	query := `
	SELECT
		p.id,
		p.title,
		p.content,
		p.created_at,
		p.updated_at,
		c.id AS category_id,
		c.name AS category_name,
		STRING_AGG(t.id::TEXT, ',') AS tag_ids,
		STRING_AGG(t.name, ',') AS tag_names
	FROM
		post p
	JOIN category c ON p.id_category = c.id
	LEFT JOIN post_tags pt ON p.id = pt.id_post
	LEFT JOIN tag t ON pt.id_tag = t.id
	WHERE c.id = $1
	GROUP BY p.id, c.id
	ORDER BY p.id;
	`

	rows, err := c.db.QueryContext(ctx, query, id)
	if err != nil {
		log.ErrorContext(ctx, "fail to query posts", "error", err)
		return nil, err
	}
	defer rows.Close()

	var posts []blog.Post

	for rows.Next() {
		var post blog.Post
		var category blog.Category
		var tagIDs, tagNames sql.NullString

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
			&category.ID,
			&category.Name,
			&tagIDs,
			&tagNames,
		)
		if err != nil {
			log.ErrorContext(ctx, "fail to scan posts", "error", err)
			return nil, err
		}

		post.Category = &category

		if tagIDs.Valid && tagIDs.String != "" {
			tagIDStrings := strings.Split(tagIDs.String, ",")
			for _, idStr := range tagIDStrings {
				tagID, err := strconv.Atoi(idStr)
				if err == nil {
					post.Tags = append(post.Tags, &blog.Tag{
						ID: tagID,
					})
				}
			}
		}

		if tagNames.Valid && tagNames.String != "" {
			tagNameStrings := strings.Split(tagNames.String, ",")

			for i, name := range tagNameStrings {
				if i < len(post.Tags) {
					post.Tags[i].Name = name
				} else {
					post.Tags = append(post.Tags, &blog.Tag{
						Name: name,
					})
				}
			}
		}
		sort.Slice(post.Tags, func(i, j int) bool {
			return post.Tags[i].ID < post.Tags[j].ID
		})

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success get all posts of category")
	return posts, nil
}

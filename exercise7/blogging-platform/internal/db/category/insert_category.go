package category

import (
	"context"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
)

func (c *Category) InsertCategory(ctx context.Context, cat blog.CategoryRequest) (*blog.Category, error) {
	log := c.logger.With("method", "InsertCategory")

	var category blog.Category
	query := `INSERT INTO category (name) VALUES ($1) RETURNING id, name, created_at, updated_at`
	err := c.db.QueryRowContext(ctx, query, cat.Name).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success insert new category")
	return &category, nil
}

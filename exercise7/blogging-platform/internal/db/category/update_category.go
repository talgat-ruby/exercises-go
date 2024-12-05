package category

import (
	"context"
	"database/sql"
	"errors"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
)

func (c *Category) UpdateCategory(ctx context.Context, id int, cat blog.CategoryRequest) (*blog.Category, error) {
	log := c.logger.With("method", "UpdateCategory")

	if cat.Name == "" {
		log.ErrorContext(ctx, "category name cannot be empty")
		return nil, errors.New("category name cannot be empty")
	}

	var category blog.Category
	query := `UPDATE category SET name = $1, updated_at = NOW() WHERE id = $2 RETURNING id, name, created_at, updated_at`
	err := c.db.QueryRowContext(ctx, query, cat.Name, id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		log.ErrorContext(ctx, "failed to update category", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success update category")
	return &category, nil
}

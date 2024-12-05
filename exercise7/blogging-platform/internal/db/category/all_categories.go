package category

import (
	"context"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
)

func (c *Category) GetAllCategories(ctx context.Context) ([]blog.Category, error) {
	log := c.logger.With("method", "GetAllCategories")

	query := `SELECT id, name, created_at, updated_at FROM category`
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		log.ErrorContext(ctx, "fail to query categories", "error", err)
		return nil, err
	}

	defer rows.Close()

	var categories []blog.Category
	for rows.Next() {
		var cat blog.Category
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.CreatedAt, &cat.UpdatedAt); err != nil {
			log.ErrorContext(ctx, "fail to scan category", "error", err)
			return nil, err
		}
		categories = append(categories, cat)
	}

	if err := rows.Err(); err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success get all categories")
	return categories, nil
}

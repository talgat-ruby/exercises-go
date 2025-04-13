package category

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
)

func (c *Category) GetInformationOfCategory(ctx context.Context, id int) (*blog.Category, error) {
	log := c.logger.With("method", "GetInformationOfCategory")

	var category blog.Category
	query := `SELECT id, name, created_at, updated_at FROM category WHERE id = $1`
	row := c.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("category with id %d not found", id)
		}
		return nil, err
	}

	log.InfoContext(ctx, "success get information of category")
	return &category, nil
}

package category

import (
	"context"
	"fmt"
)

func (c *Category) DeleteCategory(ctx context.Context, id int) error {
	log := c.logger.With("method", "DeleteCategory")

	queryDeleteCategory := `DELETE FROM category WHERE id = $1`
	deleteCategory, err := c.db.ExecContext(ctx, queryDeleteCategory, id)
	if err != nil {
		log.ErrorContext(ctx, "failed to delete category", "error", err)
		return err
	}

	countCategory, err := deleteCategory.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "failed to get rows affected", "error", err)
		return err
	}

	if countCategory == 0 {
		log.ErrorContext(ctx, "category not found", "category_id", id)
		return fmt.Errorf("category with id %d not found", id)
	}

	log.InfoContext(ctx, "success delete category")

	return nil
}

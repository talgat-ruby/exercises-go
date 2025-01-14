package expense

import (
	"context"
)

func (d *Expense) DeleteExpense(ctx context.Context, id int) error {
	log := d.logger.With("method", "DeleteExpense")

	stmt := `DELETE FROM expenses WHERE id = $1`

	row := d.db.QueryRowContext(ctx, stmt, id)
	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to query table expenses", "error", err)
		return err
	}

	return nil
}

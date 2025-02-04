package expense

import (
	"context"
)

func (d *Expense) UpdateExpense(ctx context.Context, expenseData *ModelExpense, id int) error {
	log := d.logger.With("method", "UpdateExpense")

	stmt := `UPDATE expenses 
			 SET amount = $2, description = $3
			 WHERE id = $1`

	row := d.db.QueryRowContext(ctx, stmt, id, expenseData.Amount, expenseData.Description)
	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to query table expenses", "error", err)
		return err
	}

	return nil
}

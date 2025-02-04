package expense

import (
	"context"
)

func (d *Expense) FindExpense(ctx context.Context, id int) (*ModelExpense, error) {
	log := d.logger.With("method", "FindExpenses")

	stmt := `SELECT * FROM expenses WHERE id = $1`

	row := d.db.QueryRowContext(ctx, stmt, id)
	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to query table expenses", "error", err)
		return nil, err
	}

	expense := ModelExpense{}
	if err := row.Scan(
		&expense.ID,
		&expense.Amount,
		&expense.Description,
		&expense.CreatedAt,
		&expense.UpdatedAt,
	); err != nil {
		log.ErrorContext(ctx, "fail to scan expense", "error", err)
		return nil, err
	}

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	return &expense, nil
}

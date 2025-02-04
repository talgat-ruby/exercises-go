package expense

import (
	"context"
)

func (d *Expense) CreateExpense(ctx context.Context, expenseData *CreateModelExpense) (*ModelExpense, error) {
	log := d.logger.With("method", "CreateExpenses")

	stmt := `INSERT INTO expenses (amount, description)
			 VALUES ($1, $2) 
			 RETURNING id, amount, description, created_at, updated_at`

	row := d.db.QueryRowContext(ctx, stmt, expenseData.Amount, expenseData.Description)
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

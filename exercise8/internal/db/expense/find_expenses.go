package expense

import (
	"context"
	"time"
)

func (d *Expense) FindExpenses(ctx context.Context, limit, offset int, filter string) ([]ModelExpense, error) {
	log := d.logger.With("method", "FindExpenses")

	expenses := make([]ModelExpense, 0)

	stmt := `SELECT * FROM expenses `
	var filterTime time.Time
	if filter != "" {
		stmt += `WHERE created_at > $3 `
		now := time.Now()
		switch filter {
		case "past_week":
			filterTime = now.AddDate(0, 0, -7)
			break
		case "past_month":
			filterTime = now.AddDate(0, -1, 0)
			break
		case "last_3_months":
			filterTime = now.AddDate(0, -3, 0)
			break
		}
	}
	stmt += `ORDER BY id DESC OFFSET $1 LIMIT $2`

	rows, err := d.db.QueryContext(ctx, stmt, offset, limit, filterTime)
	if err != nil {
		log.ErrorContext(ctx, "fail to query table expenses", "error", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		expense := ModelExpense{}

		if err := rows.Scan(
			&expense.ID,
			&expense.Amount,
			&expense.Description,
			&expense.CreatedAt,
			&expense.UpdatedAt,
		); err != nil {
			log.ErrorContext(ctx, "fail to scan expense", "error", err)
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	return expenses, nil
}

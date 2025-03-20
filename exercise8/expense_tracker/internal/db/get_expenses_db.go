package db

import "tracker/internal/models"

func (e *ExpencesDBSt) DBGetExpenses(id int) (models.TransactionsReponse, error) {
	transactions := models.TransactionsReponse{}
	transactions.IdUser = id

	stmt := `
	SELECT expense, comment, expense_category, created
	FROM transactions
	WHERE id_user = $1
	`
	rows, err := e.NewDb.Query(stmt, id)
	if err != nil {
		return transactions, err
	}
	defer rows.Close()
	for rows.Next() {
		transaction := models.Transaction{}
		err := rows.Scan(
			&transaction.Expense,
			&transaction.Comment,
			&transaction.ExpenseCategory,
			&transaction.CreatedAT,
		)
		if err != nil {
			return transactions, err
		}
		transactions.Transactions = append(transactions.Transactions, transaction)
	}
	if err = rows.Err(); err != nil {
		return transactions, err
	}
	return transactions, nil
}

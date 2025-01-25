package db

import (
	"log/slog"
	"tracker/internal/models"
)

func (e *ExpencesDBSt) DBEditUser(editUser models.EditUserRequest, idUser int) error {
	budgetId := 0
	tx, err := e.NewDb.Begin()
	if err != nil {
		return err
	}

	defer func() {
		p := recover()
		if p != nil {
			tx.Rollback()
		} else if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				slog.Error("failed to rollback transaction")
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				slog.Error("failed to commit transaction")
			}
		}

	}()
	stmt := `
	UPDATE budget
	SET amount= amount - $1
	WHERE id_user = $2
	RETURNING id_budget;
	`
	row := tx.QueryRow(stmt, editUser.Expense, idUser)
	err = row.Scan(
		&budgetId,
	)
	if err != nil {
		return err
	}
	query := `
	INSERT INTO transactions (id_user, id_budget, expense, comment, expense_category)
	VALUES ($1, $2, $3, $4, $5);
	`
	_, err = tx.Exec(query, idUser, budgetId, editUser.Expense, editUser.Comment, editUser.ExpenseCategory)
	if err != nil {
		return err
	}
	return nil
}

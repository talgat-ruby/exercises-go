package db

import (
	"log/slog"
	"tracker/internal/models"
)

func (e *ExpencesDBSt) NewUserDB(newUser models.NewUser) (models.NewUserResponse, error) {
	stmt := `
		INSERT INTO budget (user_id, amount)
		VALUES ($1, $2)
		RETURNING budget_id, user_id, amount
		`

	tx, err := e.NewDb.Begin()
	if err != nil {
		return models.NewUserResponse{}, err
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
	respNewUser := models.NewUserResponse{}
	row := tx.QueryRow(stmt, newUser.UserID, newUser.Amount)
	err = row.Scan(
		&respNewUser.BudgetID,
		&respNewUser.UserID,
		&respNewUser.Amount,
	)
	if err != nil {
		return models.NewUserResponse{}, err
	}
	return respNewUser, nil
}

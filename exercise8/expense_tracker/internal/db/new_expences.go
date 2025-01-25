package db

import (
	"log/slog"
	"tracker/internal/models"
)

func (e *ExpencesDBSt) NewUserDB(newUser models.NewUser, id int) (models.NewUserResponse, error) {
	stmt := `
		INSERT INTO budget (id_user, amount)
		VALUES ($1, $2)
		RETURNING id_budget, id_user, amount
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
	row := tx.QueryRow(stmt, id, newUser.Amount)
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

package authdb

import (
	"context"
	"log/slog"
	"tracker/internal/db"
	"tracker/internal/models"
)

type AuthDB interface {
	DBRegister(ctx context.Context, inp *models.DBModelUser) (*models.DBModelUser, error)
}

type authDB struct {
	database db.ExpencesDBSt
	logger   *slog.Logger
}

func NewAuthDB(db db.ExpencesDBSt) AuthDB {
	return &authDB{database: db}
}

func (h *authDB) DBRegister(ctx context.Context, inp *models.DBModelUser) (*models.DBModelUser, error) {
	log := h.logger.With("method", "Register")
	tx, err := h.database.
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to begin transaction",
			"error", err,
		)
	}
	defer func() {
		if err != nil {
			log.ErrorContext(
				ctx,
				"failed to commit transaction",
				"error", err,
			)
			tx.Rollback()

		} else {
			err = tx.Commit()
		}
	}()
	user := models.DBModelUser{}
	stmt := `INSERT INTO users (email, password_hash, salt) VALUES ($1, $2, $3) RETURNING id, email, password_hash, salt;`
	row := tx.QueryRowContext(ctx, stmt, inp.Email, inp.PasswordHash, inp.Salt)
	if err := row.Err(); err != nil {
		log.ErrorContext(
			ctx,
			"failed to insert user into database",
			"error", err,
		)
		return nil, err
	}
	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Salt,
	)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to scan user from database",
			"error", err,
		)
		return nil, err
	}

	return &user, nil

}

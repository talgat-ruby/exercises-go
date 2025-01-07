package authdb

import (
	"context"
	"errors"
	"log/slog"

	"tracker/internal/db"
	"tracker/internal/models"
)

type AuthDB interface {
	DBRegister(ctx context.Context, inp *models.DBModelUser) (*models.DBModelUser, error)
	DbLogin(email string, ctx context.Context) (*models.DBModelUser, error)
}

type AuthentificatorDB struct {
	database *db.ExpencesDBSt
	logger   *slog.Logger
	// newdb NewDb
}

func NewAuthDB(db *db.ExpencesDBSt, slogger *slog.Logger) *AuthentificatorDB {
	return &AuthentificatorDB{
		database: db,
		logger:   slogger,
	}
}

func (h *AuthentificatorDB) DBRegister(ctx context.Context, inp *models.DBModelUser) (*models.DBModelUser, error) {
	if h.database.NewDb == nil {
		return nil, errors.New("nil poinet sqldb")
	}

	log := h.logger.With("method", "Register")
	if h.logger == nil {
		return nil, errors.New("log is nil")
	}
	tx, err := h.database.NewDb.Begin()
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to begin transaction",
			"error", err,
		)
		return nil, errors.New("failed transation")
	}

	defer func() {
		p := recover()
		if p != nil {
			tx.Rollback()
			slog.Error("panic")
		} else if err != nil {
			rollbackerr := tx.Rollback()
			if rollbackerr != nil {
				log.ErrorContext(ctx, "failed to rollback transaction", "rollback_error", rollbackerr)
			}
		} else {
			commitErr := tx.Commit()
			if commitErr != nil {
				log.ErrorContext(ctx, "failed to commit transction", "commit_transaction", commitErr)
			}
		}
	}()
	user := models.DBModelUser{}
	stmt := `INSERT INTO users_ (email, password_hash, salt) VALUES ($1, $2, $3) RETURNING id, email, password_hash, salt;`
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

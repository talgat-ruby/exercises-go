package authdb

import (
	"context"
	"tracker/internal/models"
)

func (h *AuthentificatorDB) DbLogin(email string, ctx context.Context) (*models.DBModelUser, error) {
	log := h.logger.With("method", "DbLogin")
	stmt := `
		SELECT id_user, email, password_hash, salt
		FROM users_
		WHERE email = $1`
	row := h.database.NewDb.QueryRowContext(ctx, stmt, email)
	if row.Err() != nil {
		log.ErrorContext(
			ctx,
			"failed to get user",
			"error", row.Err(),
		)
		return nil, row.Err()
	}
	user := models.DBModelUser{}
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Salt,
	)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to scan user",
			"error", err,
		)
		return nil, err
	}
	log.InfoContext(ctx, "user found")
	return &user, nil
}

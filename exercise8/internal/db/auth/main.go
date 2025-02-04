package auth

import (
	"database/sql"
	"log/slog"
)

type Auth struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(logger *slog.Logger, db *sql.DB) *Auth {
	return &Auth{
		logger: logger,
		db:     db,
	}
}

type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthTokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

package auth

import (
	"log/slog"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
)

type Auth struct {
	logger *slog.Logger
	db     *db.DB
}

func New(logger *slog.Logger, db *db.DB) *Auth {
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

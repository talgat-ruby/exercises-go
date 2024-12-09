package blog

import (
	"database/sql"
	"log/slog"
)

type Blog struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(db *sql.DB, logger *slog.Logger) *Blog {
	return &Blog{
		logger: logger,
		db:     db,
	}
}

package postes

import (
	"blogging-platform/internal/db"
	"log/slog"
)

type Posts struct {
	logger *slog.Logger
	db     *db.DB
}

func New(logger *slog.Logger, db *db.DB) *Posts {
	return &Posts{
		logger: logger,
		db:     db,
	}
}

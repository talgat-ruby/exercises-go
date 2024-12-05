package tag

import (
	"database/sql"
	"log/slog"
)

type Tag struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(db *sql.DB, logger *slog.Logger) *Tag {
	return &Tag{
		logger: logger,
		db:     db,
	}
}

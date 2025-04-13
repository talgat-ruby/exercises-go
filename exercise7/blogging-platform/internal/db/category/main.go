package category

import (
	"database/sql"
	"log/slog"
)

type Category struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(db *sql.DB, logger *slog.Logger) *Category {
	return &Category{
		logger: logger,
		db:     db,
	}
}

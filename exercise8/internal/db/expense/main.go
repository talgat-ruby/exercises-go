package expense

import (
	"database/sql"
	"log/slog"
)

type Expense struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(logger *slog.Logger, db *sql.DB) *Expense {
	return &Expense{
		logger: logger,
		db:     db,
	}
}

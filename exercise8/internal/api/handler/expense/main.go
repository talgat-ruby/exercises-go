package expense

import (
	"expense_tracker/internal/db"
	"log/slog"
)

type Expense struct {
	logger *slog.Logger
	db     *db.DB
}

func New(logger *slog.Logger, db *db.DB) *Expense {
	return &Expense{
		logger: logger,
		db:     db,
	}
}

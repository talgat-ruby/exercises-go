package handler

import (
	"expense_tracker/internal/api/handler/auth"
	"expense_tracker/internal/api/handler/expense"
	"expense_tracker/internal/db"
	"log/slog"
)

type Handler struct {
	*expense.Expense
	*auth.Auth
}

func New(logger *slog.Logger, db *db.DB) *Handler {
	return &Handler{
		Expense: expense.New(logger, db),
		Auth:    auth.New(logger, db),
	}
}

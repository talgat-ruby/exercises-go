package blog

import (
	"database/sql"
	"log/slog"
)

type Blog struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(sql *sql.DB, logger *slog.Logger) *Blog {
	return &Blog{
		db:     sql,
		logger: logger,
	}
}

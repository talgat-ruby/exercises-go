package post

import (
	"database/sql"
	"log/slog"
)

type Post struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(db *sql.DB, logger *slog.Logger) *Post {
	return &Post{
		logger: logger,
		db:     db,
	}
}

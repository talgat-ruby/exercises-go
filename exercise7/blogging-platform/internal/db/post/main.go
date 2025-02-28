package post

import (
	"database/sql"
	"log/slog"
)

type Post struct {
	db     *sql.DB
	logger *slog.Logger
}

func New(db *sql.DB, logger *slog.Logger) *Post {
	return &Post{
		logger: logger,
		db:     db,
	}
}

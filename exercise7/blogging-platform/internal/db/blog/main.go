package blog

import (
	"database/sql"
	"log/slog"
)

type Blogs struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewBlog(logger *slog.Logger, db *sql.DB) *Blogs {
	return &Blogs{
		logger: logger,
		db:     db,
	}
}

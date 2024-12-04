package posts

import (
	"log/slog"

	"github.com/UAssylbek/blogging-platform/internal/db"
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

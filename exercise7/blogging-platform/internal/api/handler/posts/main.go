package posts

import (
	"log/slog"

	"github.com/webaiz/exercise7/blogging-platform/internal/db"
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

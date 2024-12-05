package tags

import (
	"log/slog"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
)

type Tags struct {
	logger *slog.Logger
	db     *db.DB
}

func New(logger *slog.Logger, db *db.DB) *Tags {
	return &Tags{
		logger: logger,
		db:     db,
	}
}

package categories

import (
	"log/slog"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
)

type Categories struct {
	logger *slog.Logger
	db     *db.DB
}

func New(logger *slog.Logger, db *db.DB) *Categories {
	return &Categories{
		logger: logger,
		db:     db,
	}
}

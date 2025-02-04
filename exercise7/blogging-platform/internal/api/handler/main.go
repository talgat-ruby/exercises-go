package handler

import (
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handler/blogs"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
	"log/slog"
)

type Handler struct {
	*blogs.Blogs
}

func NewHandler(logger *slog.Logger, db *db.ConfDB) *Handler {
	return &Handler{
		blogs.New(
			logger,
			db,
		),
	}
}

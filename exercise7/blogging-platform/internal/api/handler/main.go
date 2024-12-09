package handler

import (
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handler/blogs"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
	"log/slog"
)

type Handler struct {
	*blogs.Blogs
}

func New(logger *slog.Logger, db *db.DB) *Handler {
	return &Handler{
		Blogs: blogs.New(logger, db),
	}
}

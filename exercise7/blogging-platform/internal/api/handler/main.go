package handler

import (
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handler/posts"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
	"log/slog"
)

type Handler struct {
	*posts.Posts
}

func New(logger *slog.Logger, db *db.DB) *Handler {
	return &Handler{
		Posts: posts.New(logger, db),
	}
}

package handler

import (
	"log/slog"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handler/posts"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
)

type Handler struct {
	*posts.Posts
}

func New(logger *slog.Logger, db *db.DB) *Handler {
	return &Handler{
		Posts: posts.New(logger, db),
	}
}

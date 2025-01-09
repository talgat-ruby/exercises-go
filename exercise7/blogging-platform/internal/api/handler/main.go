package handler

import (
	"github.com/webaiz/exercise7/blogging-platform/internal/api/handler/auth"
	"github.com/webaiz/exercise7/blogging-platform/internal/api/handler/posts"
	"github.com/webaiz/exercise7/blogging-platform/internal/db"
	"log/slog"
)

type Handler struct {
	*auth.Auth
	*posts.Posts
}

func New(logger *slog.Logger, db *db.DB) *Handler {
	return &Handler{
		Auth:  auth.New(logger, db),
		Posts: posts.New(logger, db),
	}
}

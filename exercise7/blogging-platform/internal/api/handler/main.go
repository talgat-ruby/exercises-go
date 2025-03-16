package handler

import (
	"blogging-platform/internal/api/handler/posts"
	"blogging-platform/internal/db"
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

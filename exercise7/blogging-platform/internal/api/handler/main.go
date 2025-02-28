package handler

import (
	"blogging-platform/internal/api/handler/postes"
	"blogging-platform/internal/db"
	"log/slog"
)

type Handler struct {
	*postes.Posts
}

func New(logger *slog.Logger, db *db.DB) *Handler {
	return &Handler{
		Posts: postes.New(logger, db),
	}
}

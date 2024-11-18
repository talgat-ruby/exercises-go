package handler

import (
	"log/slog"
	"net/http"
)

func (h *Handler) Start(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := slog.With(
		"handler", "Start",
		"path", r.URL.Path,
	)

	// add a player to the game
	if err := h.game.Start(ctx); err != nil {
		log.ErrorContext(
			ctx,
			"game was not started",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.InfoContext(
		ctx,
		"game started",
	)
	w.WriteHeader(http.StatusNoContent)
}

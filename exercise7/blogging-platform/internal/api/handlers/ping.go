package handlers

import (
	"log/slog"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := slog.With(
		"handler", "Ping",
		"path", r.URL.Path,
	)

	log.InfoContext(
		ctx,
		"pinged",
	)
	w.WriteHeader(http.StatusOK)
}

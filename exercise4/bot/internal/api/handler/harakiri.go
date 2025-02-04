package handler

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func (h *Handler) Harakiri(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := slog.With(
		"handler", "Harakiri",
		"path", r.URL.Path,
	)
	log.InfoContext(
		ctx,
		"performing noble art of sepuku...",
	)
	_, _ = w.Write([]byte("Banzaaaaaai..."))
	go func() {
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()
}

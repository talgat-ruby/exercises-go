package handler

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Handler struct {
	name string
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := slog.With(
		"handler", "Status",
		"path", r.URL.Path,
	)
	fmt.Println(ctx, log)
	w.WriteHeader(http.StatusOK)
}

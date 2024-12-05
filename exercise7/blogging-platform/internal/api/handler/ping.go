package handler

import (
	"log/slog"
	"net/http"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	slog.Info("Received Ping request")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("pong"))

	if err != nil {
		slog.Error("failed to write response: %v", err)
	}
}

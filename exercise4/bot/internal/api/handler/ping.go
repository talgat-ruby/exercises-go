package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ping pong")
}

package handler

import (
	"log"
	"net/http"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("pong"))

	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}
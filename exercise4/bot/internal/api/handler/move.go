package handler

import (
	"encoding/json"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/dto"
	"log/slog"
	"net/http"
)

func (h *Handler) Move(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := slog.With(
		"handler", "Move",
		"path", r.URL.Path,
	)
	log.InfoContext(
		ctx,
		"move",
	)
	moveRequest := dto.RequestMove{}
	err := json.NewDecoder(r.Body).Decode(&moveRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ind := moveRequest.Board.CalculateNewPosition(moveRequest.Token)
	w.Header().Set("Content-Type", "application/json")
	moveResponse := dto.ResponseMove{Index: ind}
	err = json.NewEncoder(w).Encode(moveResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

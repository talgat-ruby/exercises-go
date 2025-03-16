package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise4/bot/pkg/httputils/request"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/ticTacToe"
)

type RequestMove struct {
	Board *ticTacToe.Board `json:"board"`
	Token ticTacToe.Token  `json:"token"`
}

type ResponseMove struct {
	Index int `json:"index"`
}

func (h *Handler) Move(w http.ResponseWriter, r *http.Request) {
	var req RequestMove
	if err := request.JSON(w, r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	position := ticTacToe.CalculateBestMove(req.Board, req.Token)
	resp := ResponseMove{Index: position}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to respond to move", http.StatusInternalServerError)
		log.Printf("Failed to encode response: %v", err)
	}
}

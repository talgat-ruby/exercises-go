package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MoveRequest struct {
	GameID string `json:"game_id"`
	Board  []int  `json:"board"`
}

type MoveResponse struct {
	Position int `json:"position"`
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "pong")
}

func handleMove(w http.ResponseWriter, r *http.Request) {
	var req MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	move := calculateMove(req.Board)

	resp := MoveResponse{Position: move}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func calculateMove(board []int) int {
	for i, val := range board {
		if val == 0 {
			return i
		}
	}
	return -1
}

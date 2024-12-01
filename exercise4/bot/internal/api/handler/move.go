package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Token string

const (
	TokenEmpty Token = " "
	TokenX     Token = "x"
	TokenO     Token = "o"
)

const (
	Cols = 3
	Rows = 3
)

type Board [Cols * Rows]Token

type RequestMove struct {
	Board Board `json:"board"`
	Token Token `json:"token"`
}

type ResponseMove struct {
	Index int `json:"index"`
}

func (h *Handler) Move(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestMove
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Received move request for token %s on board:\n%s\n", reqBody.Token, reqBody.Board.String())

	rand.Seed(time.Now().UnixNano())
	availableMoves := reqBody.Board.AvailableMoves() // Получаем все доступные ходы
	if len(availableMoves) == 0 {
		http.Error(w, "No available moves", http.StatusBadRequest)
		return
	}

	randomMove := availableMoves[rand.Intn(len(availableMoves))] // Выбираем случайный ход

	resp := ResponseMove{Index: randomMove}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

}

func (b *Board) AvailableMoves() []int {
	var moves []int

	for i := 0; i < len(b); i++ {
		if b[i] == TokenEmpty {
			moves = append(moves, i)
		}
	}

	return moves
}

func (b *Board) String() string {
	var result string
	for row := 0; row < Rows; row++ {
		result += "| "
		for col := 0; col < Cols; col++ {
			index := row*Cols + col
			result += string(b[index]) + " | "
		}
		result += "\n"
	}
	return result
}

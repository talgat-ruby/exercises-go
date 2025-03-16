package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type MoveRequest struct {
	Board [][]string `json:"board"`
}

type MoveResponse struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

func main() {
	// Read the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/move", handleMove)

	// Start the server
	log.Printf("Bot is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// handlePing handles the ping request from judge to ensure the bot is online
func handlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

// handleMove handles move requests from judge and returns a random valid move
func handleMove(w http.ResponseWriter, r *http.Request) {
	var moveReq MoveRequest
	err := json.NewDecoder(r.Body).Decode(&moveReq)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Generate a random valid move
	row, col := getRandomMove(moveReq.Board)

	moveRes := MoveResponse{
		Row: row,
		Col: col,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moveRes)
}

// getRandomMove finds an empty cell randomly and returns it as the bot's move
func getRandomMove(board [][]string) (int, int) {
	rand.Seed(time.Now().UnixNano())
	emptyCells := []struct{ row, col int }{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == "" {
				emptyCells = append(emptyCells, struct{ row, col int }{i, j})
			}
		}
	}

	if len(emptyCells) == 0 {
		return -1, -1 // No available moves
	}

	move := emptyCells[rand.Intn(len(emptyCells))]
	return move.row, move.col
}

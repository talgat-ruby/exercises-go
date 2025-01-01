package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type GameState struct {
	Board   [9]string `json:"board"`
	Current string    `json:"current"`
}

var (
	port = os.Getenv("PORT")
	url  string
)

func main() {
	if port == "" {
		log.Fatal("PORT environment variable is required")
	}

	url = fmt.Sprintf("http://localhost:%s", port)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := startPing(); err != nil {
			log.Fatalf("Failed to ping judge: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop // Wait for interrupt signal to stop the bot
}

func startPing() error {
	for {
		resp, err := http.Get(url + "/ping")
		if err != nil {
			return fmt.Errorf("ping error: %w", err)
		}
		if resp.StatusCode == http.StatusOK {
			fmt.Println("Successfully connected to the judge.")
			break
		}
		resp.Body.Close()
		time.Sleep(1 * time.Second) // Retry every second
	}

	return handleGame()
}

func handleGame() error {
	for {
		resp, err := http.Get(url + "/game")
		if err != nil {
			return fmt.Errorf("failed to get game state: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			var state GameState
			if err := json.NewDecoder(resp.Body).Decode(&state); err != nil {
				return fmt.Errorf("failed to decode game state: %w", err)
			}

			if state.Current == "" {
				// Game has ended
				fmt.Println("Game ended.")
				return nil
			}

			move, err := makeMove(state)
			if err != nil {
				return fmt.Errorf("failed to make move: %w", err)
			}

			_, err = http.Post(url+"/move", "application/json", move)
			if err != nil {
				return fmt.Errorf("failed to send move: %w", err)
			}
		}

		time.Sleep(1 * time.Second) // Poll every second
	}
}

func makeMove(state GameState) (json.RawMessage, error) {
	// Implement your logic for making a move based on the current board state
	for i, cell := range state.Board {
		if cell == "" {
			return json.Marshal(map[string]int{"move": i})
		}
	}
	return nil, fmt.Errorf("no valid moves available")
}

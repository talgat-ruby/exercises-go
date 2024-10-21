package ticTacToe

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type RequestJoin struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (p *Player) JoinGame(ctx context.Context) error {

	reqBody := RequestJoin{Name: p.Name, URL: p.URL}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal join request: %w", err)
	}

	client := &http.Client{}

	judgePort := os.Getenv("JUDGE")
	if judgePort == "" {
		judgePort = "4444"
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("http://localhost:%s/join", judgePort),
		bytes.NewBuffer(reqBodyBytes))

	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to join the game: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	log.Println("Successfully joined the game")
	return nil
}

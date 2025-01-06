package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type JoinRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func joinGame() error {
	joinURL := "http://localhost:4444/join"
	botName := "MyTicTacToeBot"
	botURL := "http://localhost:" + os.Getenv("PORT")

	reqBody := JoinRequest{
		Name: botName,
		URL:  botURL,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal join request: %w", err)
	}

	resp, err := http.Post(joinURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to send join request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to join the game, status: %d", resp.StatusCode)
	}

	log.Println("Bot successfully joined the game!")
	return nil
}

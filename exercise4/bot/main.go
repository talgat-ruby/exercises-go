package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type MoveRequest struct {
	Board []string `json:"board"`
	Token string   `json:"token"`
}

type MoveResponse struct {
	Index int `json:"index"`
}

func main() {

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	ready := startServer()
	<-ready

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {
		var req MoveRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		move := findBestMove(req.Board)

		resp := MoveResponse{Index: move}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	})

	err := registerBot(os.Getenv("NAME"), "http://localhost:"+os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Ошибка регистрации бота: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down gracefully...")
	cancel()
	time.Sleep(1 * time.Second)
}

func registerBot(name, url string) error {
	registerURL := "http://localhost:4444/join"
	botInfo := map[string]string{
		"name": name,
		"url":  url,
	}

	jsonData, err := json.Marshal(botInfo)
	if err != nil {
		return fmt.Errorf("ошибка сериализации данных бота: %w", err)
	}

	resp, err := http.Post(registerURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("ошибка отправки запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("не удалось зарегистрировать бота, статус: %d", resp.StatusCode)
	}

	log.Println("Бот успешно зарегистрирован")
	return nil
}

func findBestMove(board []string) int {
	for i, cell := range board {
		if cell == " " {
			return i
		}
	}
	return -1
}

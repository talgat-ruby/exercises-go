package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Структура для запроса на /join
type JoinRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Обработчик для Ping
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)

	ready := startServer(ctx, mux)
	<-ready

	// TODO after server start
	fmt.Println("bot server started")

	botName := os.Getenv("NAME")
	if botName == "" {
		botName = "DefaultBotName" // Можно задать имя по умолчанию
	}

	err := joinGame(botName, fmt.Sprintf("http://localhost:%s", os.Getenv("PORT")))
	if err != nil {
		fmt.Println("Failed to join game:", err)
		return
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Wait for SIGINT or SIGTERM

	fmt.Println("Shutting down bot server...")
	cancel() // Отменяем контекст, чтобы запустить завершение работы сервера
}

func joinGame(name, url string) error {
	joinRequest := JoinRequest{
		Name: name,
		URL:  url,
	}

	// Формируем JSON тело запроса
	data, err := json.Marshal(joinRequest)
	if err != nil {
		return fmt.Errorf("failed to marshal join request: %w", err)
	}

	// Отправляем POST запрос с JSON телом
	resp, err := http.Post("http://localhost:4444/join", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to join game: %w", err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Response body:", string(respBody))

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to join game, status: %s", resp.Status)
	}

	fmt.Println("Successfully joined the game")
	return nil
}

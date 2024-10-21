package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sendRequest(ctx context.Context) {
	port := os.Getenv("PORT")
	name := os.Getenv("NAME")

	// Если имя не задано в переменной окружения, установить значение по умолчанию
	if name == "" {
		name = "Unknown Player"
	}

	js, err := json.Marshal(map[string]interface{}{
		"name": name, // Используем имя из переменной окружения
		"url":  fmt.Sprintf("http://127.0.0.1:%s", port),
	})
	if err != nil {
		fmt.Println("Error during marshaling:", err)
		return
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"http://127.0.0.1:4444/join",
		bytes.NewBuffer(js),
	)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending join request:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response status:", res.StatusCode)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ready := startServer(ctx)
	sendRequest(ctx)
	<-ready

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Wait for SIGINT or SIGTERM
}

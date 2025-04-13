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
	js, err := json.Marshal(map[string]any{
		"name": name,
		"url":  fmt.Sprintf("http://127.0.0.1:%s", port),
	})
	if err != nil {
		fmt.Println("error Marshal")
		return
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/join", "http://127.0.0.1:4444"),
		bytes.NewBuffer(js),
	)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error NewRequestWithContext")
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error send join")
		return
	}
	fmt.Println(res.StatusCode)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ready := startServer(ctx)
	sendRequest(ctx)
	<-ready

	// TODO after server start

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Wait for SIGINT or SIGTERM
}

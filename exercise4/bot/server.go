package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/game/move"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// TODO: logic to response
// TODO: Don't start server if judge returns bad request

type readyListener struct {
	net.Listener
	ready chan struct{}
	once  sync.Once
}

func (l *readyListener) Accept() (net.Conn, error) {
	l.once.Do(func() { close(l.ready) })
	return l.Listener.Accept()
}

type RequestMove struct {
	Board []string `json:"board"`
	Token string   `json:"token"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Received request:", r.Method, r.URL.Path)

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}
	fmt.Println("Request Body:", string(bodyBytes))

	// Reset the request body
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var reqMove RequestMove
	if err := json.Unmarshal(bodyBytes, &reqMove); err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	board := reqMove.Board
	token := strings.ToLower(reqMove.Token) // tokens are lowercase "x" or "o"

	if len(board) != 9 {
		http.Error(w, "Invalid board size", http.StatusBadRequest)
		return
	}

	if token != "x" && token != "o" {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	index := move.GetBestMove(board, token)
	if index == -1 {
		http.Error(w, "No moves left", http.StatusBadRequest)
		return
	}

	response := map[string]int{"index": index}
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(js)
	if err != nil {
		fmt.Printf("Failed to write response: %v\n", err)

	}
}

func startServer(ctx context.Context) <-chan struct{} {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4081"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/move", moveHandler)

	server := &http.Server{
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	ready := make(chan struct{})
	list := &readyListener{Listener: listener, ready: ready}

	slog.InfoContext(
		ctx,
		"starting service",
		"port", port,
	)

	go func() {
		err := server.Serve(list)
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	return ready
}

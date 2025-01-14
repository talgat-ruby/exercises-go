package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/game"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/types"
	"io"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

type readyListener struct {
	net.Listener
	ready chan struct{}
	once  sync.Once
}

func (l *readyListener) Accept() (net.Conn, error) {
	l.once.Do(func() { close(l.ready) })
	return l.Listener.Accept()
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var moveReq types.MoveRequest
	err = json.Unmarshal(bodyBytes, &moveReq)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	cell := game.BestMove(&moveReq)
	resp := types.ResponseMove{Index: cell}

	responseData, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func startServer() <-chan struct{} {
	ready := make(chan struct{})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		panic(err)
	}

	list := &readyListener{Listener: listener, ready: ready}
	srv := &http.Server{
		IdleTimeout: 2 * time.Minute,
	}

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/move", moveHandler)

	go func() {
		err := srv.Serve(list)
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	return ready
}

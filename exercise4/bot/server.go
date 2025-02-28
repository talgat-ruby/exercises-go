package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/api/router"
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
	fmt.Fprintf(w, "pong")
}

func startServer() <-chan struct{} {
	ready := make(chan struct{})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		panic(err)
	}

	print(listener)

	// os.Getenv("NAME")

	r := router.New()

	list := &readyListener{Listener: listener, ready: ready}
	srv := &http.Server{
		Handler:     r,
		IdleTimeout: 2 * time.Minute,
	}

	go func() {
		err := srv.Serve(list)
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	return ready
}

package server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/talgat-ruby/exercises-go/exercise4/bot/handler"
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

// Start the server and handle requests using http.NewServeMux()
func StartServer() (<-chan struct{}, *http.Server) {
	ready := make(chan struct{})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	if port == "4444" {
		log.Fatalf("Fatal error: Port %s is reserved for the game.", port)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		panic(err)
	}

	list := &readyListener{Listener: listener, ready: ready}
	srv := &http.Server{
		IdleTimeout: 2 * time.Minute,
		Handler:     New(),
	}

	go func() {
		err := srv.Serve(list)
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	return ready, srv
}

func New() *http.ServeMux {
	name := os.Getenv("NAME")
	url := fmt.Sprintf("http://localhost:%s", os.Getenv("PORT"))

	han := handler.New(name, url)
	mux := http.NewServeMux()

	mux.Handle("GET /ping", http.HandlerFunc(han.Ping))
	mux.Handle("POST /move", http.HandlerFunc(han.Move))

	return mux
}

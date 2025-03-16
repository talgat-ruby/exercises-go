package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/talgat-ruby/exercises-go/exercise4/bot/server"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/ticTacToe"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ready, srv := server.StartServer()
	<-ready

	name := os.Getenv("NAME")
	if name == "" {
		name = "Player"
	}

	port := os.Getenv("PORT")
	url := fmt.Sprintf("http://localhost:%s", port)
	player := ticTacToe.New(name, url)

	if err := player.JoinGame(ctx); err != nil {
		log.Fatalf("Failed to join the game: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
}

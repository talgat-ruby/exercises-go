package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/talgat-ruby/exercises-go/exercise3/problem1/cmd/api"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/cmd/db"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/env"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	envVars := env.NewEnv()

	model := db.NewDB(envVars)
	defer model.Close()

	a := api.NewApi(model)
	a.Start(ctx, cancel, envVars)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// Wait for termination signal
	<-signalCh

	fmt.Println("gracefully shutting down...")
}

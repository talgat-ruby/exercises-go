package main

import (
	"context"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/api"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/config"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/db"
	"log/slog"
	"os"
)

func main() {
	banner, err := os.ReadFile("./resources/banner.txt")
	if err == nil {
		fmt.Println(string(banner))
	}

	ctx, cancel := context.WithCancel(context.Background())
	props, err := config.ScanDefaultProps()
	if err != nil {
		panic(err)
	}
	ctx = context.WithValue(ctx, config.PropName, props)
	configs, err := config.NewConfig(ctx)
	if err != nil {
		panic(err)
	}
	m, err := db.NewModel(configs.DB)
	if err != nil || m == nil {
		slog.Error("Failed to connect to database", err)
	}
	slog.InfoContext(ctx, "initialize service", "service", "db")

	// configure gateway service
	srv := api.NewApi(configs.Api, m)
	slog.InfoContext(ctx, "initialize service", "service", "api")
	// start gateway service
	srv.Start(ctx, cancel)
	<-ctx.Done()
	slog.InfoContext(ctx, "server was successful shutdown.")

}

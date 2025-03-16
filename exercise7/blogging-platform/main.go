package main

import (
	"blogging-platform/internal/api"
	"blogging-platform/internal/db"
	"context"
	"log/slog"
)

func main() {
	ctx := context.Background()

	db, err := db.NewDB(slog.With("service", "database"))
	if err != nil {
		panic(err)
	}

	db.Init()

	api := api.NewApi(slog.With("service", "database"), db)
	if err := api.Start(ctx); err != nil {
		panic(err)
	}
}

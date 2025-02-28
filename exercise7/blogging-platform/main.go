package main

import (
	"blogging-platform/internal/api"
	"blogging-platform/internal/db"
	"context"
	"log/slog"
)

func main() {

	ctx := context.Background()

	newDB, err := db.NewDB(slog.With("service", "database"))
	if err != nil {
		panic(err)
	}

	err = newDB.Init()
	if err != nil {
		return
	}

	newApi := api.NewApi(slog.With("service", "database"), newDB)
	if err := newApi.Start(ctx); err != nil {
		panic(err)
	}
}

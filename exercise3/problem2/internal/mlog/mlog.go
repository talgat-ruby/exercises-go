package mlog

import (
	"log"
	"log/slog"
	"os"
)

func InitLogger() {
	f, err := os.OpenFile("./logs/cli.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(f, nil)))
}

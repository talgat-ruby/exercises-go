package mlog

import (
	"log"
)

func InitLogger() {
	// fpath := "./logs/server.log"
	// out, err := os.Open(fpath)
	// if err != nil {
	// 	slog.Error("can not open log file", slog.Any("err", err))
	// 	out, err = os.Create(fpath)
	// 	if err != nil {
	// 		slog.Error("can not create log file", slog.Any("err", err))
	// 		out = os.Stdout
	// 	}
	// }
	log.Println("hello from init")

	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// slog.SetDefault(logger)
	// slog.SetLogLoggerLevel(slog.LevelInfo)
}

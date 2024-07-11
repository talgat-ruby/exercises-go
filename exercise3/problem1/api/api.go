package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/api/router"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/config"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/db"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
)

type Api struct {
	conf  *config.ApiConfig
	model *db.Model
}

func NewApi(conf *config.ApiConfig, model *db.Model) *Api {
	return &Api{
		conf:  conf,
		model: model,
	}
}

func (a *Api) Start(ctx context.Context, cancel context.CancelFunc) {
	mux := http.NewServeMux()
	router.SetupRoutes(mux, a.model)

	// start up HTTP
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.conf.Port),
		Handler: mux,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	// Listen from s different goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(ctx, "server error", "error", err)
		}

		cancel()
	}()

	slog.InfoContext(
		ctx,
		"starting api service",
		"port", a.conf.Port,
		"playground", fmt.Sprintf("http://localhost:%d/", a.conf.Port),
	)

	shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
	signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

	go func() {
		_ = <-shutdown

		slog.WarnContext(ctx, "gracefully shutting down...")
		if err := srv.Shutdown(ctx); err != nil {
			slog.ErrorContext(ctx, "server shutdown error", "error", err)
		}
	}()
}

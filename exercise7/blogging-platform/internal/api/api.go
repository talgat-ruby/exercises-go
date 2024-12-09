package api

import (
	"context"
	"errors"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handler"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/router"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/config"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
)

type Api struct {
	config *config.Config
	model  *db.DB
	logger *slog.Logger
}

func NewApi(config *config.Config, db *db.DB) *Api {
	return &Api{
		config: config,
		model:  db,
		logger: slog.With("service", "api"),
	}
}

func (a *Api) Start(ctx context.Context, cancel context.CancelFunc) {
	h := handler.NewHandler(a.logger, a.model)
	r := router.NewRouter(h)
	mux := r.Start(ctx)

	srv := &http.Server{
		Addr:        ":8080",
		Handler:     mux,
		BaseContext: func(net.Listener) context.Context { return ctx },
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(ctx, "server error", "error", err)
		}
		cancel()
	}()

	slog.InfoContext(
		ctx,
		"starting api service",
		"port", ":8080",
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

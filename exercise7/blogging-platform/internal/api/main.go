package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handler"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/router"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
)

type Api struct {
	logger *slog.Logger
	router *router.Router
	server *http.Server
}

func New(logger *slog.Logger, db *db.DB) *Api {
	h := handler.New(logger, db)
	r := router.New(h)

	return &Api{
		logger: logger,
		router: r,
	}
}

func (api *Api) Start(ctx context.Context) error {
	mux := api.router.Start(ctx)

	port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	api.server = srv

	slog.InfoContext(
		ctx,
		"starting service",
		"port", port,
	)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.ErrorContext(ctx, "service error", "error", err)
		return err
	}

	return nil
}

func (api *Api) Stop(ctx context.Context) error {
	if err := api.server.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "server shutdown error", "error", err)
		return err
	}

	return nil
}

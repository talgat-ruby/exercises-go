package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/api/router"
)

type Api struct {
	srv *http.Server
}

func New() *Api {
	return &Api{}
}

func (api *Api) Start(ctx context.Context, port string) error {
	r := router.New()

	// start up HTTP server
	api.srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	slog.InfoContext(
		ctx,
		"starting service",
		"port", port,
	)

	if err := api.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.ErrorContext(ctx, "service error", "error", err)
		return err
	}

	return nil
}

func (api *Api) Stop(ctx context.Context) error {
	if err := api.srv.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "server shutdown error", "error", err)
		return err
	}

	return nil
}

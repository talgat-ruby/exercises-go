package api

import (
	"context"
	"errors"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/api/router"
	"log/slog"
	"net"
	"net/http"
)

type Api struct {
	srv *http.Server
}

func New() *Api {
	return &Api{}
}

func (api *Api) Start(ctx context.Context, port string) error {
	r := router.New()

	api.srv = &http.Server{
		Addr:        ":" + port,
		Handler:     r,
		BaseContext: func(net.Listener) context.Context { return ctx },
	}

	slog.InfoContext(
		ctx,
		"starting bot",
		"port", port,
	)

	if err := api.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.ErrorContext(ctx, "service error", "error", err)
		return err
	}

	return nil
}

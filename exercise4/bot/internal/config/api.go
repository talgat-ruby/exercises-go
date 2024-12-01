package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/api/router"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/constants"
	"log"
	"log/slog"
	"net"
	"net/http"
)

type Api struct {
	srv  *http.Server
	port string
	ctx  context.Context
}

func (api *Api) Start(preInitFun func()) {
	r := router.New()
	ctx := context.Background()
	// start up HTTP server
	api.port = constants.PORT
	api.srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", api.port),
		Handler: r,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}
	api.ctx = ctx

	slog.InfoContext(
		api.ctx,
		"starting service",
		"port", api.port,
	)

	go preInitFun()

	if err := api.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(api.ctx, "service error", "error", err)
	}

}

func (api *Api) Stop() {
	slog.InfoContext(
		api.ctx,
		"closing service",
		"port", api.port,
	)
	if err := api.srv.Shutdown(api.ctx); err != nil {
		log.Fatal(api.ctx, "server shutdown error", "error", err)
	}
}

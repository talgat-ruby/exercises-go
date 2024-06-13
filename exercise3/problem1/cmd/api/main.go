package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise3/problem1/cmd/api/router"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/cmd/db"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/env"
)

type Api struct {
	model *db.Model
}

func NewApi(model *db.Model) *Api {
	return &Api{
		model: model,
	}
}

func (a *Api) Start(ctx context.Context, cancel context.CancelFunc, env *env.Env) {
	mux := http.NewServeMux()
	router.SetupRoutes(mux, a.model)

	// start up HTTP
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", env.AppPort),
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
}

package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handler"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/router"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
	"log/slog"
	"net"
	"net/http"
)

type Api struct {
	logger *slog.Logger
	router *router.Router
}

func NewApi(logger *slog.Logger, db *db.ConfDB) *Api {
	h := handler.NewHandler(logger, db)
	r := router.NewRouter(h)
	return &Api{
		logger: slog.With("service", "api"),
		router: r,
	}
}

func (a *Api) Start(ctx context.Context) error {
	mux := a.router.Start(ctx)

	srv := &http.Server{
		Addr:        ":8080",
		Handler:     mux,
		BaseContext: func(net.Listener) context.Context { return ctx },
	}

	fmt.Println("Staring server on :8080")
	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

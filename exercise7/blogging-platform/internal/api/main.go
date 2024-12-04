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

	"github.com/UAssylbek/blogging-platform/internal/api/handler"
	"github.com/UAssylbek/blogging-platform/internal/api/router"
	"github.com/UAssylbek/blogging-platform/internal/db"
)

type Api struct {
	logger *slog.Logger
	router *router.Router
}

func New(logger *slog.Logger, db *db.DB) *Api {
	h := handler.New(logger, db)
	r := router.New(h)

	return &Api{
		logger: logger,
		router: r,
	}
}

func (a *Api) Start(ctx context.Context) error {
	mux := a.router.Start(ctx)

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

	fmt.Printf("Starting server on :%d\n", port)
	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

package api

import (
	"context"
	"errors"
	"fmt"
	"log"
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

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	log.Printf("Start server: %d\n", port)
	err = server.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

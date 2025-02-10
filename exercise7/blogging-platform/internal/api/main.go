package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handler"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/middleware"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/router"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
)

type Api struct {
	db     *db.DB
	logger *slog.Logger
}

func New(logger *slog.Logger, db *db.DB) *Api {

	return &Api{logger: logger, db: db}
}

func (api *Api) Start(ctx context.Context) error {
	h := handler.New(slog.With("service", "handler"), api.db)
	midd := middleware.New(api.logger)
	r, err := router.New(h, midd)

	if err != nil {
		return err
	}

	mux := r.Start(ctx)

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

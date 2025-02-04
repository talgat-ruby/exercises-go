package api

import (
	"context"
	"errors"
	"expense_tracker/internal/api/handler"
	"expense_tracker/internal/api/middleware"
	"expense_tracker/internal/api/router"
	"expense_tracker/internal/db"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
)

type Api struct {
	logger *slog.Logger
	router *router.Router
}

func New(logger *slog.Logger, db *db.DB) *Api {
	midd := middleware.New(logger)
	h := handler.New(logger, db)
	r := router.New(logger, h, midd)

	return &Api{
		logger: logger,
		router: r,
	}
}

func (a *Api) Start(ctx context.Context) error {
	mux := a.router.Start(ctx)

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	fmt.Printf("Starting server on :%d\n", port)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

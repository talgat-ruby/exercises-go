package api

import (
	"blogging-platform/internal/api/handler"
	"blogging-platform/internal/api/router"
	"blogging-platform/internal/db"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Api struct {
	logger *slog.Logger
	router *router.Router
}

func NewApi(logger *slog.Logger, db *db.DB) *Api {

	handler := handler.New(logger.With("type", "handler"), db)
	router := router.New(handler)
	return &Api{
		logger: logger,
		router: router,
	}
}

func (api *Api) Start(ctx context.Context) error {

	mux := api.router.Start(ctx)
	_ = godotenv.Load()
	port, err := strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	fmt.Println("Server started")
	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

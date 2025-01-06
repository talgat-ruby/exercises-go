package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"tracker/internal/api/middleware"
	"tracker/internal/db"
	authdb "tracker/internal/db/auth_db"
	"tracker/internal/handler"
	authhandler "tracker/internal/handler/auth_handler"
	"tracker/internal/service"
)

func BasicHandlers(mux *http.ServeMux, newdb db.ExpencesDBSt) {
	serviceExpence := service.NewServiceExpence(newdb)
	handlerExpence := handler.NewHandlerExpence(serviceExpence)
	mux.Handle("POST /expenses", middleware.AuthMiddleware(http.HandlerFunc(handlerExpence.ExpensesNew)))
	mux.Handle("GET /expenses", middleware.AuthMiddleware(http.HandlerFunc(handlerExpence.ExpensesHandlerGet)))
	mux.Handle("GET /balance", middleware.AuthMiddleware(http.HandlerFunc(handlerExpence.BalanceHandler)))
}

func AuthorizationHandlers(mux *http.ServeMux, newdb *db.ExpencesDBSt) {
	logger := slog.With("service", "auth")
	authdb := authdb.NewAuthDB(newdb, logger)
	if authdb == nil {
		fmt.Println("aaaaaaaa")
		return
	}
	authHandler := authhandler.NewAuthHandler(authdb, logger)
	if authHandler == nil {
		fmt.Println("bbbbbbbbbbb")
		return
	}

	mux.HandleFunc("POST /register", authHandler.Register)
}

package api

import (
	"log"
	"log/slog"
	"net/http"

	"tracker/internal/api/middleware"
	"tracker/internal/db"
	authdb "tracker/internal/db/auth_db"
	"tracker/internal/handler"
	authhandler "tracker/internal/handler/auth_handler"
	"tracker/internal/service"
)

// type Router struct {
// 	router  *http.ServeMux

// 	midd    *middleware.Middleware
// }

func BasicHandlers(mux *http.ServeMux, newdb db.ExpencesDBSt) {
	serviceExpence := service.NewServiceExpence(newdb)
	handlerExpence := handler.NewHandlerExpence(serviceExpence)
	mux.Handle("POST /expense", middleware.AuthMiddleware(http.HandlerFunc(handlerExpence.ExpensesNew)))
	mux.Handle("GET /expenses/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlerExpence.ExpensesHandlerGet)))
	mux.Handle("GET /balance/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlerExpence.BalanceHandler)))
	mux.Handle("PUT /expense/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlerExpence.ExpenseEdit)))
}

func AuthorizationHandlers(mux *http.ServeMux, newdb *db.ExpencesDBSt) {
	logger := slog.With("service", "auth")
	authdb := authdb.NewAuthDB(newdb, logger)
	if authdb == nil {
		log.Println("nil pionter authdb")

		return
	}
	authHandler := authhandler.NewAuthHandler(authdb, logger)
	if authHandler == nil {
		log.Println("nil pionter authHandler")
		return
	}

	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)
}

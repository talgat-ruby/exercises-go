package api

import (
	"net/http"

	"tracker/internal/api/middleware"
	"tracker/internal/db"
	authdb "tracker/internal/db/auth_db"
	"tracker/internal/handler"
	authhandler "tracker/internal/handler/auth_handler"
	"tracker/internal/service"
)

func BasicHandlers(mux *http.ServeMux, newdb db.ExpencesDB) {
	serviceExpence := service.NewServiceExpence(newdb)
	handlerExpence := handler.NewHandlerExpence(serviceExpence)
	mux.Handle("POST /expenses", middleware.AuthMiddleware(http.HandlerFunc(handlerExpence.ExpensesNew)))
	mux.Handle("GET /expenses", middleware.AuthMiddleware(http.HandlerFunc(handlerExpence.ExpensesHandlerGet)))
	mux.Handle("GET /balance", middleware.AuthMiddleware(http.HandlerFunc(handlerExpence.BalanceHandler)))
}

func AuthorizationHandlers(mux *http.ServeMux, newdb db.ExpencesDB) {
	authdb := authdb.NewAuthDB(&newdb)
	authHandler := authhandler.NewAuthHandler(authdb)

	mux.HandleFunc("POST /register", authHandler.Register)
}

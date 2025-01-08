package handler

import (
	"net/http"

	"tracker/internal/service"
)

type ExpencesHandler interface {
	BalanceHandler(w http.ResponseWriter, r *http.Request)
	ExpensesHandlerGet(w http.ResponseWriter, r *http.Request)
	ExpensesNew(w http.ResponseWriter, r *http.Request)
	ExpenseEdit(w http.ResponseWriter, r *http.Request)
}

type expensesHandler struct {
	serviceExpence service.ServiceExpence
}

func NewHandlerExpence(expenceService service.ServiceExpence) ExpencesHandler {
	return &expensesHandler{serviceExpence: expenceService}
}

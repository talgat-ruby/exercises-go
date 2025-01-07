package handler

import (
	"fmt"
	"net/http"
	"tracker/internal/models"
)

func (h expensesHandler) ExpensesNew(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NewExpence")
}

func (h expensesHandler) ExpensesHandlerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Expence Get")
}

func (h expensesHandler) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("user").(*models.UserData)
	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	fmt.Println("Balance handler")
}

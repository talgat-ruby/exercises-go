package handler

import (
	"fmt"
	"net/http"
)

func (h expensesHandler) ExpensesNew(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NewExpence")
}

func (h expensesHandler) ExpensesHandlerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Expence Get")
}

func (h expensesHandler) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Balance handler")
}

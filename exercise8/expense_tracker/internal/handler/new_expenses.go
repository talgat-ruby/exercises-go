package handler

import (
	"net/http"
	"tracker/internal/models"
	"tracker/utils/request"
	"tracker/utils/respone"
)

func (h *expensesHandler) ExpensesNew(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, ok := ctx.Value("user").(*models.UserData)
	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
	newUser := &models.NewUser{}
	err := request.RequestJSON(w, r, newUser)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	responseNewUser, err := h.serviceExpence.ServiceNewUser(*newUser)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = respone.ResponseJSON(w, responseNewUser, http.StatusCreated)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

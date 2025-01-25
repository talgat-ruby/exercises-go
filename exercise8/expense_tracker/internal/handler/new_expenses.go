package handler

import (
	"log/slog"
	"net/http"
	"tracker/internal/models"
	"tracker/utils/request"
	"tracker/utils/respone"
)

func (h *expensesHandler) ExpensesNew(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, ok := ctx.Value("user").(*models.UserData)
	if !ok {
		slog.Error("Context", "error", "user not found")
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	authHeader := r.Header.Get("Authorization")
	authToken := authHeader[len("Bearer "):]

	newUser := &models.RequestNewUser{}
	err := request.RequestJSON(w, r, newUser)
	if err != nil {
		slog.Error("RequestJSON", "error", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	responseNewUser, err := h.serviceExpence.ServiceNewUser(*&newUser.NewUserRequest, authToken)
	if err != nil {
		slog.Error("ServiceNewUser", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = respone.ResponseJSON(w, responseNewUser, http.StatusCreated)
	if err != nil {
		slog.Error("ResponseJSON", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

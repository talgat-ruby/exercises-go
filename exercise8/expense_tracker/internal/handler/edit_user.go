package handler

import (
	"log/slog"
	"net/http"
	"tracker/internal/models"
	"tracker/utils/request"
	"tracker/utils/respone"
)

func (h *expensesHandler) ExpenseEdit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, ok := ctx.Value("user").(*models.UserData)
	if !ok {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
	editUser := &models.EditUser{}
	err := request.RequestJSON(w, r, editUser)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		slog.Error("failed to write", "error", err)
		return
	}
	err = h.serviceExpence.ServiceEditUser(*editUser)
	if err != nil {
		http.Error(w, "status internal server", http.StatusInternalServerError)
		slog.Error("failed to write", "error", err)
		return
	}
	err = respone.ResponseJSON(w, nil, http.StatusNoContent)
	if err != nil {
		http.Error(w, "status internal server", http.StatusInternalServerError)
		slog.Error("failed to write", "error", err)
		return
	}
	slog.Info("succes update expense")

}

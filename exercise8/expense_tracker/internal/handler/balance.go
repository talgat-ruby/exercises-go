package handler

import (
	"log/slog"
	"net/http"
	"strconv"
	"tracker/internal/models"
	"tracker/utils/respone"
)

func (h *expensesHandler) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("user").(*models.UserData)
	if !ok {
		slog.Error("User not found")
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("incorrect id", "error", err)
		http.Error(w, "incorrect id", http.StatusBadRequest)
		return
	}
	balance, err := h.serviceExpence.ServiceBalance(idInt)
	if err != nil {
		slog.Error("internal server error", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	respone.ResponseJSON(w, balance, http.StatusOK)
	slog.Info("succes balance")
}

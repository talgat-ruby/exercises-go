package handler

import (
	"log/slog"
	"net/http"
	"strconv"
	"tracker/internal/models"
	"tracker/utils/respone"
)

func (h *expensesHandler) ExpensesHandlerGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, ok := ctx.Value("user").(*models.UserData)
	if !ok {
		slog.Error("unauthorized user")
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
	transactions, err := h.serviceExpence.ServiceGetExpenses(idInt)
	if err != nil {
		slog.Error("inernal server error", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	err = respone.ResponseJSON(w, transactions, http.StatusOK)
	if err != nil {
		slog.Error("error write json", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	slog.Info("succes get expense")
}

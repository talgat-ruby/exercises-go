package expense

import (
	"expense_tracker/internal/db/expense"
	"expense_tracker/pkg/httputils/response"
	"net/http"
	"strconv"
)

type DeleteExpenseResponse struct {
	Data *expense.ModelExpense `json:"data"`
}

func (h *Expense) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("DeleteExpense")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed parse query offset",
			"error", err,
		)
		http.Error(w, "invalid query offset", http.StatusBadRequest)
		return
	}

	err = h.db.DeleteExpense(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := response.JSON(
		w,
		http.StatusOK,
		nil,
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		return
	}

	log.InfoContext(ctx, "success deleted expense")
	return
}

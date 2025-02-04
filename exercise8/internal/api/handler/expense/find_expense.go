package expense

import (
	"expense_tracker/internal/db/expense"
	"expense_tracker/pkg/httputils/response"
	"net/http"
	"strconv"
)

type FindExpenseResponse struct {
	Data *expense.ModelExpense `json:"data"`
}

func (h *Expense) FindExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("FindExpense")

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

	dbData, err := h.db.FindExpense(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	resp := FindExpenseResponse{
		Data: dbData,
	}

	if err := response.JSON(
		w,
		http.StatusOK,
		resp,
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		return
	}

	log.InfoContext(ctx, "success find expense", "successfully get expense")
	return
}

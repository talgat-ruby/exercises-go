package expense

import (
	"expense_tracker/internal/db/expense"
	"expense_tracker/pkg/httputils/request"
	"expense_tracker/pkg/httputils/response"
	"net/http"
	"strconv"
)

type UpdateExpenseResponse struct {
	Data *expense.ModelExpense `json:"data"`
}

func (h *Expense) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("UpdateExpense")

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

	requestBody := &UpdateExpenseResponse{}
	if err := request.JSON(w, r, requestBody); err != nil {
		log.ErrorContext(ctx, "failed to parse request body", "error", err)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	err = h.db.UpdateExpense(ctx, requestBody.Data, id)
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

	log.InfoContext(ctx, "success create expense")
	return
}

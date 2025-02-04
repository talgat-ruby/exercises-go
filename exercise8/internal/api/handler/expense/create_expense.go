package expense

import (
	"expense_tracker/internal/db/expense"
	"expense_tracker/pkg/httputils/request"
	"expense_tracker/pkg/httputils/response"
	"net/http"
)

type CreateExpenseResponse struct {
	Data *expense.CreateModelExpense `json:"data"`
}

func (h *Expense) CreateExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("CreateExpense")

	requestBody := &CreateExpenseResponse{}
	if err := request.JSON(w, r, requestBody); err != nil {
		log.ErrorContext(ctx, "failed to parse request body", "error", err)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	dbResp, err := h.db.CreateExpense(ctx, requestBody.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	resp := FindExpenseResponse{
		Data: dbResp,
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

	log.InfoContext(ctx, "success create expense")
	return
}

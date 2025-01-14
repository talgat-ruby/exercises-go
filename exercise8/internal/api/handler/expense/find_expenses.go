package expense

import (
	"expense_tracker/internal/db/expense"
	"expense_tracker/pkg/httputils/response"
	"net/http"
	"strconv"
	"time"
)

type FindExpensesResponse struct {
	Data []expense.ModelExpense `json:"data"`
}

func (h *Expense) FindExpenses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("FindExpenses")

	query := r.URL.Query()
	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed parse query offset",
			"error", err,
		)
		http.Error(w, "invalid query offset", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed parse query limit",
			"error", err,
		)
		http.Error(w, "invalid query limit", http.StatusBadRequest)
		return
	}

	filter := query.Get("filter")
	if filter == "" {
		filter = time.Time{}.String()
	}

	dbData, err := h.db.FindExpenses(ctx, limit, offset, filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	resp := FindExpensesResponse{
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

	log.InfoContext(ctx, "success find expenses", "number of expenses", len(resp.Data))
	return
}

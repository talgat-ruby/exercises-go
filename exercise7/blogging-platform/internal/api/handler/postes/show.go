package postes

import (
	"blogging-platform/pkg/httputils/response"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
)

func (h *Posts) PostsShow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	h.logger.With("method", "PostsShow")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		h.logger.ErrorContext(ctx, "id to int conv error", "error", err)
	}
	res, err := h.db.PostsShow(ctx, int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not found", http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := response.JSON(
		w,
		http.StatusOK,
		res,
	); err != nil {
		h.logger.ErrorContext(ctx, "Error json response", "error", err)
	}
}

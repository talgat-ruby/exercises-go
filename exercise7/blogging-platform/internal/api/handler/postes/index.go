package postes

import (
	"blogging-platform/pkg/httputils/response"
	"net/http"
)

func (h *Posts) PostsIndex(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	h.logger.With("method", "PostsIndex")

	term := r.URL.Query().Get("term")
	h.logger.ErrorContext(ctx, "s string", "term", term)

	res, err := h.db.PostsIndex(ctx, term)
	if err != nil {
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

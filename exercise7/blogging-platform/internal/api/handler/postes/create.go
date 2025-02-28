package postes

import (
	"blogging-platform/internal/db/post"
	"blogging-platform/pkg/httputils/request"
	"blogging-platform/pkg/httputils/response"
	"net/http"
)

func (h *Posts) PostsCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	h.logger.With("method", "PostsCreate")

	data := &post.Model{}
	if err := request.JSON(w, r, data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := h.db.PostsCreate(ctx, data)
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

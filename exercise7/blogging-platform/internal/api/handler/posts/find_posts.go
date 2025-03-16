package posts

import (
	"net/http"
	"strconv"

	"github.com/webaiz/exercise7/blogging-platform/internal/db/post"
	"github.com/webaiz/exercise7/blogging-platform/pkg/httputils/response"
)

type FindPostsResponse struct {
	Data []post.ModelPost `json:"data"`
}

func (h *Posts) FindPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("method", "FindPosts")

	query := r.URL.Query()
	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"fail parse query offset",
			"error", err,
		)
		http.Error(w, "invalid query offset", http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"fail parse query limit",
			"error", err,
		)
		http.Error(w, "invalid query limit", http.StatusBadRequest)
		return
	}

	dbResp, err := h.db.FindPosts(ctx, offset, limit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := FindPostsResponse{
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

	log.InfoContext(
		ctx,
		"success find posts",
		"number_of_posts", len(resp.Data),
	)
	return
}

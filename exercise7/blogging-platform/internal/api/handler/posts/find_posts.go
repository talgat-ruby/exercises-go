package posts

import (
	"net/http"

	"github.com/UAssylbek/blogging-platform/internal/db/post"
	"github.com/UAssylbek/blogging-platform/pkg/httputils/response"
)

type FindPostsResponse struct {
	Data []post.ModelPost `json:"data"`
}

func (h *Posts) FindPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := h.logger.With("method", "FindPosts")

	dbResp, err := h.db.FindPosts(ctx)

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

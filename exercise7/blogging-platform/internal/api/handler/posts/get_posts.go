package posts

import (
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

type GetPostsResponse struct {
	Data []models.Blog `json:"data"`
}

func (h *Posts) GetPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "GetPosts")

	res, err := h.db.DBGetPosts(ctx)
	if err != nil {
		log.ErrorContext(
			ctx,
			"data base",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := GetPostsResponse{
		Data: res,
	}
	err = response.JSON(w, http.StatusOK, resp)
	if err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		return
	}
	log.InfoContext(
		ctx,
		"succes find posts",
		"number of posts", len(resp.Data),
	)
	return
}

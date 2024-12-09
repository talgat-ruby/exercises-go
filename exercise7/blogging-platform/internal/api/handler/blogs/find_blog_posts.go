package blogs

import (
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
	"net/http"
	"strconv"
)

type FindBlogPostsResponse struct {
	Data []blog.BlogPost `json:"data"`
}

func (h *Blogs) FindBlogPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "FindBlogPosts")

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

	dbResp, err := h.db.FindBlogPosts(ctx, offset, limit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := FindBlogPostsResponse{
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
		"success find movies",
		"number_of_movies", len(resp.Data),
	)
	return
}

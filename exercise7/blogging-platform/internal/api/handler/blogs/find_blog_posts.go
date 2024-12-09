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

	offset := 0
	limit := 10

	query := r.URL.Query()

	if offsetStr := query.Get("offset"); offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err != nil {
			log.ErrorContext(ctx, "invalid offset parameter", "error", err)
			http.Error(w, "invalid offset parameter", http.StatusBadRequest)
			return
		}
		if parsedOffset < 0 {
			http.Error(w, "offset cannot be negative", http.StatusBadRequest)
			return
		}
		offset = parsedOffset
	}

	if limitStr := query.Get("limit"); limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil {
			log.ErrorContext(ctx, "invalid limit parameter", "error", err)
			http.Error(w, "invalid limit parameter", http.StatusBadRequest)
			return
		}
		if parsedLimit <= 0 {
			http.Error(w, "limit must be positive", http.StatusBadRequest)
			return
		}
		if parsedLimit > 100 { // Maximum limit
			parsedLimit = 100
		}
		limit = parsedLimit
	}

	dbResp, err := h.db.FindBlogPosts(ctx, offset, limit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := FindBlogPostsResponse{
		Data: dbResp,
	}

	if err := response.JSON(w, http.StatusOK, resp); err != nil {
		log.ErrorContext(ctx, "failed to write JSON response", "error", err)
		return
	}

	log.InfoContext(
		ctx,
		"success find movies",
		"number_of_movies", len(resp.Data),
	)
	return
}

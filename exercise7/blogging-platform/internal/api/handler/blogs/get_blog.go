package blogs

import (
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
	"net/http"
	"strconv"
)

type GetBlogResponse struct {
	Data *blog.ModelBlogs `json:"data"`
}

func (h *Blogs) GetBlog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "GetBlog")

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to convert id to int",
			"error", err,
		)
		http.Error(w, "failed to convert id to int", http.StatusBadRequest)
		return
	}
	dbResp, err := h.db.GetBlog(ctx, int64(id))

	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := GetBlogResponse{
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
		"success find blog",
		"blog id", resp.Data.ID,
	)
	return
}

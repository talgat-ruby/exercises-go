package blogs

import (
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
	"net/http"
	"strconv"
)

type DeleteBlogRequest struct {
	Data *blog.BlogModel `json:"data"`
}

type DeleteBlogResponse struct {
	Data *blog.BlogModel `json:"data"`
}

func (b *Blogs) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	b.logger.Info("method", "Delete blog")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		b.logger.Error(
			"failed to convert id to int",
			"error", err,
		)
		http.Error(w, "failed to convert id to int", http.StatusBadRequest)
		return
	}

	err = b.db.DeleteBlog(id)
	if err != nil {
		b.logger.Error(
			"failed to query from db",
			"error", err,
		)
		http.Error(w, "failed to query from db", http.StatusInternalServerError)
		return
	}

	response.JSON(
		w,
		http.StatusOK,
		"successfully deleted",
	)
	return
}

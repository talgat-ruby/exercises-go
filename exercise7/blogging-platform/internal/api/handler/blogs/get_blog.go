package blogs

import (
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
	"net/http"
	"strconv"
)

func (b *Blogs) GetBlogById(w http.ResponseWriter, r *http.Request) {
	b.logger.Debug("method", "Find blog")
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

	blog, err := b.db.GetBlogById(id)
	if err != nil {
		b.logger.Error(
			"Blog not found by id",
			"error", err,
		)
		http.Error(w, "Blog not found by id", http.StatusBadRequest)
		return
	}
	response.JSON(
		w,
		http.StatusOK,
		blog,
	)
	return
}

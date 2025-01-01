package blogs

import (
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
	"net/http"
)

func (b *Blogs) GetBlogs(w http.ResponseWriter, r *http.Request) {
	_ = b.logger.With("method", "Find blog")
	offset := 0
	limit := 10
	blogs, err := b.db.GetBlogs(offset, limit)
	if err != nil {
		b.logger.Error(
			"failed to convert id to int",
			"error", err,
		)
		http.Error(w, "failed to convert id to int", http.StatusBadRequest)
		return
	}
	response.JSON(
		w,
		http.StatusOK,
		blogs,
	)
	return
}

package blogs

import (
	"encoding/json"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
	"net/http"
)

type CreateBlogRequest struct {
	Data *blog.BlogModel `json:"data"`
}

type CreateBlogResponse struct {
	Data *blog.BlogModel `json:"data"`
}

func (b *Blogs) AddBlog(w http.ResponseWriter, r *http.Request) {
	b.logger.Debug("method", "Create blog")

	decoder := json.NewDecoder(r.Body)
	rBody := blog.BlogModel{}
	err := decoder.Decode(&rBody)
	if err != nil {
		b.logger.Error("Error in request body")
		return
	}

	createdBlog, err := b.db.AddBlog(&rBody)
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
		createdBlog,
	)
	return
}

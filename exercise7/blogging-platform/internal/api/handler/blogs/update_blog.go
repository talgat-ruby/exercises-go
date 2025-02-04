package blogs

import (
	"encoding/json"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
	"net/http"
	"strconv"
)

type UpdateBlogRequest struct {
	Data *blog.BlogModel `json:"data"`
}

type UpdateBlogResponse struct {
	Data *blog.BlogModel `json:"data"`
}

func (b *Blogs) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	b.logger.Debug("method", "Update blog")
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
	decoder := json.NewDecoder(r.Body)
	rBody := blog.BlogModel{}
	err = decoder.Decode(&rBody)
	if err != nil {
		b.logger.Error("Error in request body")
		return
	}

	err = b.db.UpdateBlog(id, &rBody)
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
		"successfully updated",
	)
	return
}

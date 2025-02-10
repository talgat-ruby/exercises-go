package blogs

import (
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
	"net/http"
)

type GetBlogsResponse struct {
	Data []blog.ModelBlogs `json:"data"`
}

func (b *Blogs) GetBlogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := b.logger.With("method", "GetBlogs")

	dbResp, err := b.db.GetBlogs(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := GetBlogsResponse{
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
		"success find blogs",
		"number_of_blogs", len(resp.Data),
	)
	return

}

package blogs

import (
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/auth"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/request"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
	"net/http"
)

type CreateBlogRequest struct {
	Data *blog.ModelBlogs `json:"data"`
}

type CreateBlogResponse struct {
	Data *blog.ModelBlogs `json:"data"`
}

func (h *Blogs) CreateBlog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "CreateBlog")

	user, ok := ctx.Value("user").(*auth.UserData)
	if !ok {
		log.ErrorContext(
			ctx,
			"failed to type cast user data",
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("user: %+v\n", *user)

	// request parse
	requestBody := &CreateBlogRequest{}

	if err := request.JSON(w, r, requestBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			"error", err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	// db request
	dbResp, err := h.db.CreateBlog(ctx, requestBody.Data)

	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, "failed to query from db", http.StatusInternalServerError)
		return
	}

	if dbResp == nil {
		log.ErrorContext(
			ctx,
			"row is empty",
		)
		http.Error(w, "row is empty", http.StatusInternalServerError)
		return
	}

	// response
	resp := CreateBlogResponse{
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
		"success insert blog",
		"blog id", resp.Data.ID,
	)
	return
}

package handler

import (
	"log/slog"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handler/categories"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handler/posts"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handler/tags"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
)

type Handler struct {
	*categories.Categories
	*posts.Posts
	*tags.Tags
}

func New(logger *slog.Logger, db *db.DB) *Handler {
	return &Handler{
		Categories: categories.New(logger, db),
		Posts:      posts.New(logger, db),
		Tags:       tags.New(logger, db),
	}
}

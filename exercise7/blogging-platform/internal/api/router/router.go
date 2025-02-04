package router

import (
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handlers"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/service"
)

func SetupRouter(service *service.PostService) *http.ServeMux {
	mux := http.NewServeMux()

	handler := handlers.New(service)

	mux.Handle("GET /ping", http.HandlerFunc(handlers.Ping))
	mux.Handle("GET /posts", http.HandlerFunc(handler.GetPosts))
	mux.Handle("POST /posts", http.HandlerFunc(handler.CreatePost))
	mux.Handle("PUT /posts/", http.HandlerFunc(handler.UpdatePost))
	mux.Handle("DELETE /posts/", http.HandlerFunc(handler.DeletePost))

	return mux
}

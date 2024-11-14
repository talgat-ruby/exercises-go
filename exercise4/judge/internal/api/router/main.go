package router

import (
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/api/handler"
)

func New() *http.ServeMux {
	han := handler.New()
	mux := http.NewServeMux()

	mux.Handle("GET /status", http.HandlerFunc(han.Status))
	mux.Handle("POST /join", http.HandlerFunc(han.Join))

	// Only admin can start game, start it manually
	mux.Handle("POST /start", http.HandlerFunc(han.Start))

	return mux
}

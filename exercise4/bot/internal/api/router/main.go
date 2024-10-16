package router

import (
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/api/handler"
	"net/http"
)

func New() *http.ServeMux {
	han := handler.New()
	mux := http.NewServeMux()

	mux.Handle("GET /player1/ping", http.HandlerFunc(han.Status))
	return mux
}

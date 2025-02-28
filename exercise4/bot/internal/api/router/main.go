package router

import (
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/api/handler"
)

func New() *http.ServeMux {
	han := handler.New()
	mux := http.NewServeMux()

	mux.Handle("GET /ping", http.HandlerFunc(han.Ping))
	mux.Handle("POST /move", http.HandlerFunc(han.Move))

	return mux
}

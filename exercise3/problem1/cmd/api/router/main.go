package router

import (
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise3/problem1/cmd/db"
)

type ErrJson struct {
	Message string `json:"message"`
}

func SetupRoutes(mux *http.ServeMux, model *db.Model) {
	todoH := &todoHandler{
		model: model,
	}

	mux.HandleFunc("GET /todo", todoH.GetTodos)
}

package router

import (
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/db"
	"net/http"
)

type ErrJson struct {
	Message string `json:"message"`
}

func SetupRoutes(mux *http.ServeMux, model *db.Model) {
	handler := &todoHandler{
		model: model,
	}

	mux.HandleFunc("GET /todo/{todoId}", handler.GetTodoHandler)
	mux.HandleFunc("GET /todo", handler.GetAllTodoHandler)
	mux.HandleFunc("POST /todo", handler.CreateTodoHandler)
	mux.HandleFunc("PATCH /todo/{todoId}", handler.PatchTodoHandler)
	mux.HandleFunc("DELETE /todo/{todoId}", handler.DeleteTodoHandler)
}

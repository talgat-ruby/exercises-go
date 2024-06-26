package router

import (
	"net/http"

	"exercise3/problem1/internal/db"
)

type ErrJson struct {
	Message string `json:"message"`
}

func SetupRoutes(mux *http.ServeMux, model *db.Model) {
	todoH := &todoHandler{
		model: model,
	}

	mux.HandleFunc("GET /todo/{id}", todoH.GetTodoHandler)
	mux.HandleFunc("GET /todo", todoH.GetTodosHandler)
	mux.HandleFunc("POST /todo", todoH.CreateTodoHandler)
	mux.HandleFunc("PATCH /todo/{id}", todoH.UpdateTodoHandler)
	mux.HandleFunc("DELETE /todo/{id}", todoH.DeleteTodoHandler)
}

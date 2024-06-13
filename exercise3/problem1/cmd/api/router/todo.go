package router

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise3/problem1/cmd/db"
)

type todoHandler struct {
	model *db.Model
}

type TodoListItem struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (h *todoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	slog.Info("Get todos")

	emplsModel, err := h.model.SelectTodos()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	todos := make([]*TodoListItem, 0, len(emplsModel))
	for _, todo := range emplsModel {
		todos = append(todos, &TodoListItem{
			ID:          todo.ID,
			Description: todo.Description,
			Done:        todo.Done,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

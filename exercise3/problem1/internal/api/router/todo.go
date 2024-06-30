package router

import (
	"encoding/json"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/db"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type todoHandler struct {
	model *db.Model
}

type Todo struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (h *todoHandler) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

	slog.Info("Get employee with id: ", strId)

	id, err := strconv.Atoi(strId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	todoModel, err := h.model.SelectTodo(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	todo := &Todo{
		ID:          todoModel.ID,
		Description: todoModel.Description,
		Done:        todoModel.Done,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *todoHandler) GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Get employees")

	todosModel, err := h.model.SelectTodos()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	todos := make([]*Todo, len(todosModel))
	for i, todo := range todosModel {
		todos[i] = &Todo{
			ID:          todo.ID,
			Description: todo.Description,
			Done:        todo.Done,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *todoHandler) CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	todo := &Todo{}

	if r.Header.Get("Content-Type") == "application/json" {
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	todoModel := &db.TodoModel{
		Description: todo.Description,
		Done:        todo.Done,
	}

	if err := h.model.InsertTodo(todoModel); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

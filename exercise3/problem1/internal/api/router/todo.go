package router

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"exercise3/problem1/internal/db"
)

type todoHandler struct {
	model *db.Model
}

type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	IsDone      bool   `json:"done"`
}

func (h *todoHandler) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

	slog.Info("Get todo with id: ", strId)

	id, err := strconv.Atoi(strId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	todoModel, err := h.model.GetTodo(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	if todoModel == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrJson{Message: "Record not found"})
		return
	}

	todo := &Todo{
		ID:          todoModel.ID,
		Description: todoModel.Description,
		IsDone:      todoModel.Done,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *todoHandler) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Get todos")

	todosModel, err := h.model.GetTodos()
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
			IsDone:      todo.Done,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *todoHandler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	todo := &Todo{}

	if r.Header.Get("Content-Type") == "application/json" {
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	todoModel := &db.TodoModel{
		Description: todo.Description,
	}

	if err := h.model.CreateTodo(todoModel); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *todoHandler) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

	slog.Info("todo with id to update: ", strId)

	id, err := strconv.Atoi(strId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	var data struct {
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.model.UpdateTodo(data.Description, id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *todoHandler) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

	slog.Info("todo with id to delete: ", strId)

	id, err := strconv.Atoi(strId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	if err := h.model.DeleteTodo(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrJson{Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

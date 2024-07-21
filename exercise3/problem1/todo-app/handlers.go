package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// getTodos возвращает все задачи
func getTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := executeQuery("SELECT id, description, done, created_at, updated_at FROM todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// addTodo добавляет новую задачу
func addTodo(w http.ResponseWriter, r *http.Request) {
	var todoCreate TodoCreate
	if err := json.NewDecoder(r.Body).Decode(&todoCreate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := executeQuery("INSERT INTO todos (description, done, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())", todoCreate.Description, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// getTodoByID возвращает задачу по ID
func getTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["todoId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	row := executeQueryRow("SELECT id, description, done, created_at, updated_at FROM todos WHERE id = $1", todoID)
	var todo Todo
	if err := row.Scan(&todo.ID, &todo.Description, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Todo not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// updateTodoByID обновляет задачу по ID
func updateTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["todoId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var todoUpdate TodoUpdate
	if err := json.NewDecoder(r.Body).Decode(&todoUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = executeQuery("UPDATE todos SET description = $1, done = $2, updated_at = NOW() WHERE id = $3", todoUpdate.Description, todoUpdate.Done, todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// deleteTodoByID удаляет задачу по ID
func deleteTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["todoId"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = executeQuery("DELETE FROM todos WHERE id = $1", todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

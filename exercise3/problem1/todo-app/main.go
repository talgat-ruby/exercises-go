package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Инициализация базы данных
	err := initDB()
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/todo", getTodos).Methods("GET")
	router.HandleFunc("/todo", addTodo).Methods("POST")
	router.HandleFunc("/todo/{todoId}", getTodoByID).Methods("GET")
	router.HandleFunc("/todo/{todoId}", updateTodoByID).Methods("PATCH")
	router.HandleFunc("/todo/{todoId}", deleteTodoByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

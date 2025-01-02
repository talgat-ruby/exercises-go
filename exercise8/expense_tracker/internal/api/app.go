package api

import (
	"log"
	"net/http"

	"tracker/internal/db"
)

func StartServer() error {
	mux := http.NewServeMux()
	err, newdb := db.Init()
	if err != nil {
		return err
	}
	dbExpence := db.NewExpenceDB(newdb)
	BasicHandlers(mux, dbExpence)
	log.Println("Server is running on port 8080")
	port := ":8080"
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	return server.ListenAndServe()
}

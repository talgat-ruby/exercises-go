package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"tracker/internal/db"
)

func StartServer(dbExpence *db.ExpencesDBSt) error {
	mux := http.NewServeMux()
	AuthorizationHandlers(mux, dbExpence)
	BasicHandlers(mux, *dbExpence)
	log.Println("Server is running on port 8080")
	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	return server.ListenAndServe()
}

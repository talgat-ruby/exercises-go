package router

import (
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/db"
	"net/http"
)

type ErrJson struct {
	Message string `json:"message"`
}

func SetupRoutes(mux *http.ServeMux, model *db.Model) {
	emplH := &employeeHandler{
		model: model,
	}

	mux.HandleFunc("GET /ping", GetPingHandler)
	mux.HandleFunc("GET /employee/{id}", emplH.GetEmployeeHandler)
	mux.HandleFunc("GET /employee", emplH.GetEmployeesHandler)
	mux.HandleFunc("POST /employee", emplH.CreateEmployeeHandler)
}

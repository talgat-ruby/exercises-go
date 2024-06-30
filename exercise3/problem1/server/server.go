package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise3/problem1/api/router"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/config"
)

type Server struct {
	AppConfig *config.Config
	Router    *router.Router
}

func NewServer(appConfig *config.Config, router *router.Router) *Server {
	slog.Info("hello from new server")

	return &Server{
		AppConfig: appConfig,
		Router:    router,
	}
}

func (server *Server) Run() {
	router := server.Router.InitRoutes()
	fmt.Println("ListenAndServe")
	err := http.ListenAndServe(":"+server.AppConfig.Port, router)
	if err != nil {
		return
	}
}

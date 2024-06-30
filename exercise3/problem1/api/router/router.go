package router

import (
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise3/problem1/api/controllers/todo"
)

type Router struct {
	TODORoutes *todo.ControllerRoute
}

func NewRouter(todoRoutes *todo.ControllerRoute) *Router {
	slog.Info("hello from new Router")

	return &Router{
		TODORoutes: todoRoutes,
	}
}

func (routes *Router) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router = routes.TODORoutes.Route(router)
	return router
}

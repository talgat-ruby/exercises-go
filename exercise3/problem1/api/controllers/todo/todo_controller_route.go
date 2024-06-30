package todo

import (
	"log/slog"
	"net/http"
)

type ControllerRoute struct {
	Controller *Controller
}

func NewControllerRoute(controller *Controller) *ControllerRoute {
	slog.Info("hello from new Controller route")

	return &ControllerRoute{
		Controller: controller,
	}
}

func (route *ControllerRoute) Route(router *http.ServeMux) *http.ServeMux {
	slog.Info("create new routes")
	router.HandleFunc("GET /todo", route.Controller.GetAllTasks)
	router.HandleFunc("GET /todo/{id}", route.Controller.GetTask)
	router.HandleFunc("POST /todo", route.Controller.CreateTask)
	router.HandleFunc("PATCH /todo/{id}", route.Controller.UpdateTask)
	router.HandleFunc("DELETE /todo/{id}", route.Controller.DeleteTask)
	return router
}

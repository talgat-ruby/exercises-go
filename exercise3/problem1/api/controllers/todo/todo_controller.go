package todo

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	todoModels "github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/models/todo"
	todoServices "github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/services/todo"
)

type Controller struct {
	service *todoServices.Service
}

// NewController - мето для регистрации контроллера DI
func NewController(service *todoServices.Service) *Controller {
	slog.Info("hello from new Controller")
	return &Controller{
		service: service,
	}
}

func (controller *Controller) CreateTask(w http.ResponseWriter, r *http.Request) {
	request := new(todoModels.CreateModels)
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&request)
	if strings.TrimSpace(request.Description) == "" {
		slog.Error("invalid input")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	task, err := controller.service.Create(request)
	if err != nil {
		slog.Error("Controller.CreateTask:can't create task", slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(task)
	if err != nil {
		slog.Error("can't encode task record", slog.Any("err", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (controller *Controller) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((r.PathValue("id")))
	if err != nil {
		slog.Error("can't get id from pathValue", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request := new(todoModels.UpdateModels)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	if err != nil {
		slog.Error("error when decode request", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(request.Description) == "" {
		slog.Error("invalid input")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	task, err := controller.service.Update(id, request)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		slog.Error("error when update", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(task)
	if err != nil {
		slog.Error("error when Marshal", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (controller *Controller) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((r.PathValue("id")))
	if err != nil {
		slog.Error("can't get id from pathValue", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task, err := controller.service.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		slog.Error("can't get task from service", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(task)
	if err != nil {
		slog.Error("can't encode task record", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (controller *Controller) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := controller.service.GetAll()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.Write([]byte("[]"))
			w.WriteHeader(http.StatusOK)
			return
		}

		slog.Error("can't get task from service", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		slog.Error("can't encode task record", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (controller *Controller) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((r.PathValue("id")))
	if err != nil {
		slog.Error("can't get id from pathValue", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = controller.service.Delete(id)
	if err != nil {
		slog.Error("error when delete task record", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

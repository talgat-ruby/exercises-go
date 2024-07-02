package router

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/db"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type todoHandler struct {
	model *db.Model
}

type TodoResponse struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created"`
	UpdatedAt   time.Time `json:"updated"`
}

type createTodoDto struct {
	Description string `json:"description"`
}

const (
	invalidId           = "Invalid ID supplied"
	todoId              = "todoId"
	todoNotFound        = "Todo not found"
	validationException = "Validation exception"
	contentType         = "Content-Type"
	appJson             = "application/json"
)

func (handler *todoHandler) GetTodoHandler(response http.ResponseWriter, request *http.Request) {
	strId := request.PathValue(todoId)
	id, err := strconv.Atoi(strId)
	if err != nil || id < 1 {
		response.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(response).Encode(ErrJson{Message: invalidId})
		return
	}

	todoModel, err := handler.model.SelectTodo(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(response).Encode(ErrJson{Message: todoNotFound})
		} else {
			response.WriteHeader(http.StatusUnprocessableEntity)
			_ = json.NewEncoder(response).Encode(ErrJson{Message: validationException})
		}
		return
	}
	todoR := &TodoResponse{
		Id:          todoModel.Id,
		Description: todoModel.Description,
		Done:        todoModel.Done,
		CreatedAt:   todoModel.CreatedAt,
		UpdatedAt:   todoModel.UpdatedAt,
	}
	response.Header().Set(contentType, appJson)
	err = json.NewEncoder(response).Encode(todoR)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error in GetTodoByIdHandler - " + err.Error())
		return
	}
}

func (handler *todoHandler) GetAllTodoHandler(response http.ResponseWriter, request *http.Request) {
	todoModels, err := handler.model.FindAllTodo()
	if err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(response).Encode(ErrJson{Message: validationException})
		return
	}
	todos := make([]*TodoResponse, len(todoModels))
	for id, todoModel := range todoModels {
		todoR := &TodoResponse{
			Id:          todoModel.Id,
			Description: todoModel.Description,
			Done:        todoModel.Done,
			CreatedAt:   todoModel.CreatedAt,
			UpdatedAt:   todoModel.UpdatedAt,
		}
		todos[id] = todoR
	}

	response.Header().Set(contentType, appJson)
	err = json.NewEncoder(response).Encode(todos)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error in GetAllTodoHandler - " + err.Error())
		return
	}
}

func (handler *todoHandler) CreateTodoHandler(response http.ResponseWriter, request *http.Request) {
	if request.Header.Get(contentType) != appJson {
		response.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(response).Encode(ErrJson{Message: "Content-Type must be application/json"})
		return
	}

	descDto := createTodoDto{}
	err := json.NewDecoder(request.Body).Decode(&descDto)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(response).Encode(ErrJson{Message: "Invalid input"})
		slog.Error("Error in CreateTodoHandler - " + err.Error())
		return
	}
	err = handler.model.CreateTodo(descDto.Description)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(response).Encode("Successfully created")
}

func (handler *todoHandler) PatchTodoHandler(response http.ResponseWriter, request *http.Request) {
	if request.Header.Get(contentType) != appJson {
		response.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(response).Encode(ErrJson{Message: "Content-Type must be application/json"})
		return
	}
	strId := request.PathValue(todoId)
	id, err := strconv.Atoi(strId)
	if err != nil || id < 1 {
		response.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(response).Encode(ErrJson{Message: invalidId})
		return
	}
	todoModel, err := handler.model.SelectTodo(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(response).Encode(ErrJson{Message: todoNotFound})
		} else {
			response.WriteHeader(http.StatusUnprocessableEntity)
			_ = json.NewEncoder(response).Encode(ErrJson{Message: validationException})
		}
		return
	}

	updatedTodo, errMapper := requestBodyTodoMapper(request, todoModel)
	if errMapper != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(response).Encode(ErrJson{Message: validationException})
	}

	if updatedTodo == nil {
		response.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(response).Encode("Successfully updated")
		return
	}

	errUpdate := handler.model.UpdateTodo(updatedTodo)
	if errUpdate != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(response).Encode(ErrJson{Message: validationException})
		return
	}

	response.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(response).Encode("Successfully updated")
}

func (handler *todoHandler) DeleteTodoHandler(response http.ResponseWriter, request *http.Request) {
	strId := request.PathValue(todoId)
	id, err := strconv.Atoi(strId)
	if err != nil || id < 1 {
		response.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(response).Encode(ErrJson{Message: invalidId})
		return
	}

	_, errSelect := handler.model.SelectTodo(id)
	if errSelect != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(response).Encode(ErrJson{Message: todoNotFound})
		} else {
			response.WriteHeader(http.StatusUnprocessableEntity)
			_ = json.NewEncoder(response).Encode(ErrJson{Message: validationException})
		}
		return
	}

	errDelete := handler.model.DeleteTodo(id)
	if errDelete != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(response).Encode(ErrJson{Message: validationException})
		return
	}
	response.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(response).Encode("Successfully deleted")
}

func requestBodyTodoMapper(request *http.Request, todoModel *db.TodoModel) (*db.TodoModel, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		slog.Error("Error in requestBodyTodoMapper - " + err.Error())
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(request.Body)
	var todoMap map[string]interface{}
	err = json.Unmarshal(body, &todoMap)
	if err != nil {
		slog.Error("Error in requestBodyTodoMapper - " + err.Error())
		return nil, err
	}
	if todoMap == nil {
		return nil, nil
	}
	if val, ok := todoMap["description"]; ok {
		todoModel.Description = val.(string)
	}
	if val, ok := todoMap["done"]; ok {
		todoModel.Done = val.(bool)
	}
	return todoModel, nil
}

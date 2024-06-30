package todo

import (
	"log/slog"

	todoEntities "github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/entities/todo"
	todoModels "github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/models/todo"

	"github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/repositories/todo"
)

type Service struct {
	repository todo.Repository
}

func NewService(repository todo.Repository) *Service {
	slog.Info("hello from new Service")
	return &Service{repository: repository}
}

func (service *Service) Create(model *todoModels.CreateModels) (*todoEntities.Entity, error) {
	task := &todoEntities.Entity{
		Description: model.Description,
	}
	err := service.repository.Create(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (service *Service) GetById(id int) (*todoEntities.Entity, error) {
	return service.repository.GetById(id)
}

func (service *Service) GetAll() ([]*todoEntities.Entity, error) {
	return service.repository.GetAll()
}

func (service *Service) Update(id int, updateModel *todoModels.UpdateModels) (*todoEntities.Entity, error) {
	task := &todoEntities.Entity{
		Id:          id,
		Description: updateModel.Description,
		Done:        updateModel.Done,
	}

	err := service.repository.Update(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (service *Service) Delete(id int) error {
	return service.repository.Delete(id)
}

package todo

import (
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/entities/todo"
)

type Repository interface {
	Create(*todo.Entity) error
	Update(*todo.Entity) error
	Delete(int) error
	GetById(int) (*todo.Entity, error)
	GetAll() ([]*todo.Entity, error)
}

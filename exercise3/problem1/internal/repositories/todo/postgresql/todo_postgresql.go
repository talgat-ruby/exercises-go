package postgresql

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"

	todoEntities "github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/entities/todo"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/repositories/todo"
)

type Repository struct {
	db *sql.DB
}

// Create implements todo.Repository.
func (r *Repository) Create(task *todoEntities.Entity) error {
	query := "INSERT INTO todo(description) VALUES ($1) RETURNING id, done, create_at, update_at"
	return r.db.QueryRow(query, task.Description).Scan(&task.Id, &task.Done, &task.CreateAt, &task.UpdateAt)
}

// Delete implements todo.Repository.
func (r *Repository) Delete(id int) error {
	query := "DELETE FROM todo WHERE id=$1"
	_, err := r.db.Exec(query, id)
	return err
}

// GetAll implements todo.Repository.
func (r *Repository) GetAll() ([]*todoEntities.Entity, error) {
	query := "SELECT id, description, done, create_at, update_at FROM todo"
	//.Scan(&exercise.Description, &exercise.Done, &exercise.CreateAt, &exercise.UpdateAt)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]*todoEntities.Entity, 0)
	for rows.Next() {
		task := todoEntities.Entity{}
		err = rows.Scan(&task.Id, &task.Description, &task.Done, &task.CreateAt, &task.UpdateAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

// GetById implements todo.Repository.
func (r *Repository) GetById(id int) (*todoEntities.Entity, error) {
	task := todoEntities.Entity{
		Id: id,
	}
	query := "SELECT description, done, create_at, update_at FROM todo WHERE id=$1"
	slog.Info("id is", slog.Int("id", id))
	slog.Info("get by id postgres")
	err := r.db.QueryRow(query, id).Scan(&task.Description, &task.Done, &task.CreateAt, &task.UpdateAt)
	slog.Info("after get by id postgres")

	if err != nil {
		return nil, err
	}

	return &task, nil
}

// Update implements todo.Repository.
func (r *Repository) Update(task *todoEntities.Entity) error {
	query := "UPDATE todo SET description=$1, done=$2, update_at=CURRENT_TIMESTAMP WHERE id=$3 RETURNING create_at, update_at"
	return r.db.QueryRow(query, task.Description, task.Done, task.Id).Scan(&task.CreateAt, &task.UpdateAt)
}

func NewRepository(db *sql.DB) todo.Repository {
	return &Repository{db: db}
}

package db

import (
	"database/sql"
	"errors"
	"log/slog"
	"time"
)

type TodoModel struct {
	Id          int
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (model *Model) SelectTodo(id int) (*TodoModel, error) {
	statement := "SELECT id, description, done, created_at, updated_at FROM todo WHERE id=$1"
	row := model.db.QueryRow(statement, id)
	todo := &TodoModel{}

	if err := row.Scan(&todo.Id, &todo.Description, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Info("Todo does with id %v not found", id)
		} else {
			slog.Error("Error in SelectTodo - " + err.Error())
		}
		return nil, err
	}
	return todo, nil
}

func (model *Model) FindAllTodo() ([]*TodoModel, error) {
	errMsg := "Error in FindAllTodo - "
	statement := "SELECT id, description, done, created_at, updated_at FROM todo"
	rows, err := model.db.Query(statement)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			slog.Error(errMsg + err.Error())
			panic(err)
		}
	}(rows)

	if err != nil {
		slog.Error(errMsg + err.Error())
		return nil, err
	}

	var todos []*TodoModel
	for rows.Next() {
		todo := &TodoModel{}
		if err := rows.Scan(&todo.Id, &todo.Description, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			slog.Error(errMsg + err.Error())
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (model *Model) CreateTodo(description string) error {
	statement := "insert into todo (description) values ($1)"
	_, err := model.db.Exec(statement, description)

	if err != nil {
		slog.Error("Error in CreateTodo - " + err.Error())
		return err
	}
	return nil
}

func (model *Model) UpdateTodo(update *TodoModel) error {
	statement := "update todo set description=$1, done=$2, updated_at=$3 where id=$4"
	_, err := model.db.Exec(statement, update.Description, update.Done, update.UpdatedAt, update.Id)
	if err != nil {
		slog.Error("Error in UpdateTodo - " + err.Error())
		return err
	}
	return nil
}

func (model *Model) DeleteTodo(id int) error {
	statement := "delete from todo where id=$1"
	_, err := model.db.Exec(statement, id)
	if err != nil {
		slog.Error("Error in DeleteTodo - " + err.Error())
		return err
	}
	return nil
}

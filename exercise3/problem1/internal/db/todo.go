package db

import (
	"database/sql"
	"errors"
)

type TodoModel struct {
	ID          int
	Description string
	Done        bool
}

func (m *Model) GetTodo(id int) (*TodoModel, error) {
	statement := "SELECT id, description, done FROM todos WHERE id=$1"
	row := m.db.QueryRow(statement, id)
	todo := &TodoModel{}

	if err := row.Scan(&todo.ID, &todo.Description, &todo.Done); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return todo, nil
}

func (m *Model) GetTodos() ([]*TodoModel, error) {
	statement := "SELECT id, description, done FROM todos"
	rows, err := m.db.Query(statement)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	todos := make([]*TodoModel, 0)

	for rows.Next() {
		todo := &TodoModel{}
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Done); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (m *Model) CreateTodo(todo *TodoModel) error {
	statement := "INSERT INTO todos (description)  VALUES ($1)"
	_, err := m.db.Exec(statement, todo.Description)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) UpdateTodo(description string, id int) error {
	statement := "UPDATE todos SET description = $1 WHERE id = $2"
	_, err := m.db.Exec(statement, description, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) DeleteTodo(id int) error {
	statement := "DELETE FROM todos  WHERE id = $1"
	_, err := m.db.Exec(statement, id)
	if err != nil {
		return err
	}

	return nil
}

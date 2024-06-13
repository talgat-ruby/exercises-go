package db

import (
	"database/sql"
)

type TodoModel struct {
	ID          int
	Description string
	Done        bool
}

func (m *Model) SelectTodos() ([]*TodoModel, error) {
	statement := "SELECT id, description, done FROM todo"
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

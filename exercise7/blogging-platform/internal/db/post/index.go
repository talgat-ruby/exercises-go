package post

import (
	"context"
	"database/sql"
)

func (m *Post) PostsIndex(ctx context.Context, term string) ([]Model, error) {
	log := m.logger.With("post_method", "PostsIndex")

	var rows *sql.Rows
	var err error

	models := make([]Model, 0)

	stmt := `SELECT id, title, content, category, tags, created_at, updated_at FROM posts`

	if term != "" {
		stmt += ` WHERE title ILIKE $1 OR content ILIKE $1 OR category ILIKE $1`
		rows, err = m.db.QueryContext(ctx, stmt, "%"+term+"%")
	} else {
		rows, err = m.db.QueryContext(ctx, stmt)
	}

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		model := Model{}
		if err := rows.Scan(&model.ID, &model.Title, &model.Content, &model.Category, &model.Tags, &model.CreatedAt, &model.UpdatedAt); err != nil {
			log.ErrorContext(ctx, "failed index from db", "error", err)
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}

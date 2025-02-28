package post

import "context"

func (m *Post) PostsShow(ctx context.Context, id int64) (*Model, error) {
	log := m.logger.With("post_method", "PostsShow")

	stmt := `SELECT id, title, content, category, tags, created_at, updated_at FROM posts WHERE id = $1`

	row := m.db.QueryRowContext(ctx, stmt, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	model := Model{}

	if err := row.Scan(&model.ID, &model.Title, &model.Content, &model.Category, &model.Tags, &model.CreatedAt, &model.UpdatedAt); err != nil {
		log.ErrorContext(ctx, "failed show of db", "error", err)
		return nil, err
	}

	return &model, nil
}

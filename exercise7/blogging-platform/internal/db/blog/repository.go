package blog

import "fmt"

func (b *Blogs) GetBlogs(offset, limit int) ([]BlogModel, error) {
	query := `SELECT * FROM blogs OFFSET $1 LIMIT $2`

	rows, err := b.db.Query(query, offset, limit)
	if err != nil {
		b.logger.Error("fail to query table movie", "error", err)
		return nil, err
	}

	blogs := make([]BlogModel, 0)
	defer rows.Close()

	for rows.Next() {
		blog := BlogModel{}
		if err := rows.Scan(
			&blog.ID,
			&blog.Name,
			&blog.Description,
			&blog.CreatedAt,
			&blog.UpdatedAt,
		); err != nil {
			b.logger.Error("fail to query table movie", "error", err)
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := rows.Err(); err != nil {
		b.logger.Error("fail to scan rows", "error", err)
		return nil, err
	}

	return blogs, nil
}

func (b *Blogs) GetBlogById(id int) (*BlogModel, error) {
	query := `SELECT * FROM blogs WHERE id = $1`

	rows := b.db.QueryRow(query, id)
	if err := rows.Err(); err != nil {
		b.logger.Error("fail to query table movie", "error", err)
		return nil, err
	}

	blogs := BlogModel{}
	if err := rows.Scan(
		&blogs.ID,
		&blogs.Name,
		&blogs.Description,
		&blogs.CreatedAt,
		&blogs.UpdatedAt,
	); err != nil {
		b.logger.Error("fail to query table movie", "error", err)
		return nil, err
	}
	return &blogs, nil
}

func (b *Blogs) AddBlog(data *BlogModel) (*BlogModel, error) {
	query := `INSERT INTO blogs (name, description) VALUES ($1, $2) RETURNING id, name, description, created_at, updated_at`

	rows := b.db.QueryRow(query, data.Name, data.Description)
	if err := rows.Err(); err != nil {
		b.logger.Error("fail to query table movie", "error", err)
		return nil, err
	}
	blogs := BlogModel{}
	if err := rows.Scan(
		&blogs.ID,
		&blogs.Name,
		&blogs.Description,
		&blogs.CreatedAt,
		&blogs.UpdatedAt,
	); err != nil {
		b.logger.Error("fail to query table movie", "error", err)
		return nil, err
	}
	return &blogs, nil
}

func (b *Blogs) UpdateBlog(id int, data *BlogModel) error {
	query := `UPDATE blogs SET name = $2, description = $3 WHERE id = $1`

	resp, err := b.db.Exec(query, id, data.Name, data.Description)
	if err != nil {
		b.logger.Error("fail to query table movie", "error", err)
		return err
	}

	num, err := resp.RowsAffected()
	if err != nil {
		b.logger.Error("failed to RowsAffected")
		return fmt.Errorf("failed to RowsAffected")
	}
	if num == 0 {
		b.logger.Error("blogs not found", "id", id)
		return fmt.Errorf("blogs not found id = %d", id)
	}
	b.logger.Info("blogs successfully updated")
	return nil
}

func (b *Blogs) DeleteBlog(id int) error {
	query := `DELETE FROM blogs WHERE id = $1`

	resp, err := b.db.Exec(query, id)
	if err != nil {
		b.logger.Error("fail to query table movie", "error", err)
		return err
	}

	num, err := resp.RowsAffected()
	if err != nil {
		b.logger.Error("failed to RowsAffected")
		return fmt.Errorf("failed to RowsAffected")
	}
	if num == 0 {
		b.logger.Error("blogs not found", "id", id)
		return fmt.Errorf("blogs not found id = %d", id)
	}
	b.logger.Info("blogs successfully deleted")
	return nil
}

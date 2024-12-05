package tag

import (
	"context"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
)

func (t *Tag) GetAllTags(ctx context.Context) ([]blog.Tag, error) {
	log := t.logger.With("method", "GetAllTags")

	query := `SELECT id, name, created_at, updated_at FROM tag`
	rows, err := t.db.QueryContext(ctx, query)
	if err != nil {
		log.ErrorContext(ctx, "fail to query tags", "error", err)
		return nil, err
	}

	defer rows.Close()

	var tags []blog.Tag
	for rows.Next() {
		var tag_one blog.Tag
		if err := rows.Scan(&tag_one.ID, &tag_one.Name, &tag_one.CreatedAt, &tag_one.UpdatedAt); err != nil {
			log.ErrorContext(ctx, "fail to scan tags", "error", err)
			return nil, err
		}
		tags = append(tags, tag_one)
	}

	if err := rows.Err(); err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success get all tags")
	return tags, nil
}

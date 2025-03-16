package post

import (
	"context"
	"database/sql"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
)

func (p *Post) InsertPost(ctx context.Context, req blog.PostRequest) (*blog.Post, error) {
	log := p.logger.With("method", "InsertPost")

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		log.ErrorContext(ctx, "fail to begin transaction", "error", err)
		return nil, err
	}
	defer tx.Rollback()

	id_category, err := p.GetCategoryIDByName(ctx, tx, req.Category)
	if err != nil {
		log.ErrorContext(ctx, "failed to get category by name", "error", err)
		return nil, err
	}

	category := &blog.Category{
		ID:   id_category,
		Name: req.Category,
	}

	if id_category == 0 {
		category, err = p.InsertCat(ctx, tx, *category)
		if err != nil {
			log.ErrorContext(ctx, "failed to create category", "error", err)
			return nil, err
		}
	}

	id_category = category.ID
	post, err := p.InsertInfoPost(ctx, tx, id_category, req)
	if err != nil {
		log.ErrorContext(ctx, "failed to create post", "error", err)
		return nil, err
	}
	post.Category = category

	tags, err := p.InsertTags(ctx, tx, post.ID, req.Tags)
	if err != nil {
		log.ErrorContext(ctx, "failed to insert tags", "error", err)
		return nil, err
	}
	post.Tags = tags

	if err := tx.Commit(); err != nil {
		log.ErrorContext(ctx, "fail to commit transaction", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success insert new post")
	return post, nil
}

func (p *Post) InsertInfoPost(ctx context.Context, tx *sql.Tx, id_category int, req blog.PostRequest) (*blog.Post, error) {
	log := p.logger.With("method", "InsertPost")
	var post blog.Post
	query := `INSERT INTO post (id_category,title,"content") VALUES ($1, $2, $3) RETURNING id, title, content, created_at, updated_at`
	err := tx.QueryRowContext(ctx, query, id_category, req.Title, req.Content).Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	return &post, nil
}

func (p *Post) GetCategoryIDByName(ctx context.Context, tx *sql.Tx, name string) (int, error) {
	log := p.logger.With("method", "InsertPost")
	query := `SELECT category.id FROM category WHERE name= $1`
	var id int
	err := tx.QueryRowContext(ctx, query, name).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		log.ErrorContext(ctx, "failed to get category by name", "error", err)
		return 0, err
	}

	return id, nil
}

func (p *Post) InsertCat(ctx context.Context, tx *sql.Tx, cat blog.Category) (*blog.Category, error) {
	log := p.logger.With("method", "InsertPost")
	var category blog.Category
	query := `INSERT INTO category (name) VALUES ($1) RETURNING id, name, created_at, updated_at`
	err := tx.QueryRowContext(ctx, query, cat.Name).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success insert new category")
	return &category, nil
}

func (p *Post) InsertTags(ctx context.Context, tx *sql.Tx, id_post int, tags []string) ([]*blog.Tag, error) {
	log := p.logger.With("method", "InsertPost")
	var insertedTags []*blog.Tag

	for _, tagName := range tags {
		var tag blog.Tag

		err := tx.QueryRowContext(ctx, "SELECT id, name, created_at, updated_at FROM tag WHERE name = $1", tagName).Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt)

		if err == sql.ErrNoRows {

			err := tx.QueryRowContext(ctx, "INSERT INTO tag(name) VALUES($1) RETURNING id, name, created_at, updated_at", tagName).Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt)
			if err != nil {
				log.ErrorContext(ctx, "fail to insert tag", "error", err)
				return nil, err
			}
		} else if err != nil {
			log.ErrorContext(ctx, "fail to check if tag exists", "error", err)
			return nil, err
		}

		insertedTags = append(insertedTags, &tag)

		err = p.InsertPostTags(ctx, tx, tag.ID, id_post)
		if err != nil {
			log.ErrorContext(ctx, "fail to insert tag-post relation", "error", err)
			return nil, err
		}
	}

	return insertedTags, nil
}

func (p *Post) InsertPostTags(ctx context.Context, tx *sql.Tx, tagID, postID int) error {
	log := p.logger.With("method", "InsertPost")
	var count int
	err := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM post_tags WHERE id_post = $1 AND id_tag = $2", postID, tagID).Scan(&count)
	if err != nil {
		log.ErrorContext(ctx, "fail to check if post-tag relation exists", "error", err)
		return err
	}

	if count == 0 {
		_, err := tx.ExecContext(ctx, "INSERT INTO post_tags(id_post, id_tag) VALUES($1, $2)", postID, tagID)
		if err != nil {
			log.ErrorContext(ctx, "fail to insert post-tag relation", "error", err)
			return err
		}
	}

	return nil
}

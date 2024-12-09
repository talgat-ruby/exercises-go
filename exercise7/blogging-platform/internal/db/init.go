package db

import (
	"context"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
)

func (db *DB) Init(ctx context.Context) error {
	log := db.logger.With("method", "Init")

	// Create blog_posts table
	stmt := blog.BlogPostSchema

	if _, err := db.pg.ExecContext(ctx, stmt); err != nil {
		log.ErrorContext(ctx, "failed to create blog_posts table", "error", err)
		return fmt.Errorf("creating blog_posts table: %w", err)
	}

	// Seed initial blogs posts
	seedStmt := `
INSERT INTO blog_posts (title, content, category, tags)
VALUES 
    (
        'Getting Started with Go',
        'Go is a statically typed, compiled programming language designed at Google...',
        'Programming',
        ARRAY['Go', 'Programming', 'Tutorial']
    ),
    (
        'Understanding PostgreSQL Arrays',
        'PostgreSQL provides native array support, allowing you to store multiple values...',
        'Database',
        ARRAY['PostgreSQL', 'Database', 'Arrays']
    ),
    (
        'Web Development Best Practices',
        'When building web applications, it''s important to follow these best practices...',
        'Web Development',
        ARRAY['Web', 'Development', 'Best Practices']
    )
ON CONFLICT (id) DO NOTHING;  -- Prevents duplicate inserts on repeated runs
`

	if _, err := db.pg.ExecContext(ctx, seedStmt); err != nil {
		log.ErrorContext(ctx, "failed to seed blog_posts table", "error", err)
		return fmt.Errorf("seeding blog_posts table: %w", err)
	}

	log.InfoContext(ctx, "successfully initialized blog_posts table")
	return nil
}

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
	const SeedStatement = `
WITH 
-- First, insert blog posts
blog_posts_insert AS (
    INSERT INTO blog_posts (title, content, category)
    VALUES 
        (
            'Getting Started with Go',
            'Go is a statically typed, compiled programming language designed at Google...',
            'Programming'
        ),
        (
            'Understanding PostgreSQL Arrays',
            'PostgreSQL provides native array support, allowing you to store multiple values...',
            'Database'
        ),
        (
            'Web Development Best Practices',
            'When building web applications, it''s important to follow these best practices...',
            'Web Development'
        )
    ON CONFLICT (id) DO NOTHING
    RETURNING id
),
-- Then, insert tags
tags_insert AS (
    INSERT INTO tags (name)
    VALUES 
        ('Go'),
        ('Programming'),
        ('Tutorial'),
        ('PostgreSQL'),
        ('Database'),
        ('Arrays'),
        ('Web'),
        ('Development'),
        ('Best Practices')
    ON CONFLICT (name) DO NOTHING
    RETURNING id, name
)
-- Finally, create relationships
INSERT INTO blog_posts_tags (blog_post_id, tag_id)
SELECT bp.id, t.id
FROM blog_posts_insert bp
CROSS JOIN tags_insert t
WHERE 
    (bp.id = 1 AND t.name IN ('Go', 'Programming', 'Tutorial'))
    OR (bp.id = 2 AND t.name IN ('PostgreSQL', 'Database', 'Arrays'))
    OR (bp.id = 3 AND t.name IN ('Web', 'Development', 'Best Practices'))
ON CONFLICT DO NOTHING;
`
	if _, err := db.pg.ExecContext(ctx, SeedStatement); err != nil {
		log.ErrorContext(ctx, "failed to seed blog_posts table", "error", err)
		return fmt.Errorf("seeding blog_posts table: %w", err)
	}

	log.InfoContext(ctx, "successfully initialized blog_posts table")
	return nil
}

SELECT EXISTS (
    SELECT 1
    FROM posts
    WHERE schemaname = 'public'
    AND tablename = 'posts'
);

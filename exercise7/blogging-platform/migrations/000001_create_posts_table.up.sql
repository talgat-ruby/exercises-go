CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT now(),
    title text NOT NULL,
    content text NOT NULL,
    category text NOT NULL,
    tags text[] NOT NULL
)
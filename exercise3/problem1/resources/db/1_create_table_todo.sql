CREATE TABLE IF NOT EXISTS todo (
    id  SERIAL PRIMARY KEY,
    description VARCHAR,
    done BOOLEAN default false,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp default current_timestamp
);
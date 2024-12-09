CREATE TABLE IF NOT EXISTS blogs (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR NOT NULL,
                                     description TEXT NOT NULL,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
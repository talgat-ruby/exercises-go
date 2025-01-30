*

-- Create enums
CREATE TYPE weight_unit AS ENUM ('kg', 'lbs');
CREATE TYPE distance_unit AS ENUM ('km', 'miles');
CREATE TYPE user_role AS ENUM ('user', 'admin');

-- Users table with simplified fields
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(100) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    salt TEXT NOT NULL,
    role user_role DEFAULT 'user',
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- Master table of exercises
CREATE TABLE exercise_data (
    exercise_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Workout plans
CREATE TABLE workout_plans (
    workout_plan_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Exercise logs
CREATE TABLE exercise_logs (
    exercise_log_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    workout_plan_id INTEGER,
    exercise_id INTEGER NOT NULL,
    workout_date DATE NOT NULL,
    workout_time TIME,
    set_number INTEGER NOT NULL,
    repetitions INTEGER NOT NULL,
    weight DECIMAL(5,2),
    weight_unit weight_unit,
    distance DECIMAL(5,2),
    distance_unit distance_unit,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE,
    FOREIGN KEY (exercise_id)
        REFERENCES exercise_data(exercise_id)
        ON DELETE RESTRICT,
    FOREIGN KEY (workout_plan_id)
        REFERENCES workout_plans(workout_plan_id)
        ON DELETE SET NULL
);

-- Create indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_workout_plans_user ON workout_plans(user_id);
CREATE INDEX idx_exercise_logs_user ON exercise_logs(user_id);
CREATE INDEX idx_exercise_logs_date ON exercise_logs(workout_date);
*/
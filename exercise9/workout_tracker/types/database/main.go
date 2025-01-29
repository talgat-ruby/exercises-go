package database

/*
1. user data, workout plan, exercise data

create table workout_plans(
	workout_plan_id serial  primary key,
	workout_date DATE NOT NULL,
	workout_time DATE NOT NULL,
	isActive BOOLEAN NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

create table exercise_data(
	exerciseId serial  primary key,
	name_exercise varchar(255) unique not null,
	description text,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

(not on the assignment) dont inlcude it yet, do it on next iteration and ugrade of the project.
create table workout_tags(
	id serial primary key,
	name varchar(255),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

create table comments_on_workout(
	comments_on_workout_id serial,
	workout_plan_id int,
	name_exercise varchar(255) unique not null,
	description text,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	Primary key(comments_on_workout_id),
	Constraint fk_workout
		FOREIGN KEY(workout_plan_id)
			REFERENCES workout_plans(workout_plan_id)
);

*/

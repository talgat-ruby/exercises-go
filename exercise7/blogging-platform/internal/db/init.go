package db

func (db *DB) Init() error {
	stmt := `CREATE TABLE IF NOT EXISTS posts (
		id	SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		content TEXT DEFAULT NULL,
		category TEXT DEFAULT NULL,
		tags TEXT[],
		created_at TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NOW(),
		updated_at TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT NOW()
	)`

	if _, err := db.pg.Exec(stmt); err != nil {
		return err
	}

	return nil
}

package db

type DB struct {
}

func New() (*DB, error) {
	return &DB{}, nil
}

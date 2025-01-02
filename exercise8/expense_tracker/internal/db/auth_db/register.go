package authdb

import (
	"tracker/internal/db"
)

type AuthDB interface{}

type authDB struct {
	database *db.ExpencesDB
}

func NewAuthDB(db *db.ExpencesDB) AuthDB {
	return &authDB{database: db}
}

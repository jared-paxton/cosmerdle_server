package db

import "database/sql"

type DatabaseAccessor interface{}

type databaseAccessor struct {
	db *sql.DB
}

func NewDatabaseAccessor(db *sql.DB) DatabaseAccessor {
	return &databaseAccessor{
		db: db,
	}
}

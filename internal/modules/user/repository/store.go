package repository

import "database/sql"

type Store interface {
	Queries
	// TODO
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

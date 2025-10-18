package db

import "database/sql"

// Store provides all functions to execute db queries and transactions.
// It is an interface that our handler will use, making it easy to mock for tests.
// By embedding *Queries, our Store interface will automatically have all the methods
// that sqlc generated for us (e.g., CreateUser).
type Store interface {
	Querier
	//TODO: add methods for transcations here in the future
}

// SQLStore provides all functions to execute SQL queries and transactions.
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

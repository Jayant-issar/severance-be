package handler

import "github.com/Jayant-issar/severance-backend/internal/database/db"

//Returns a a new handler and store dependecies of all the handler function

type Handler struct {
	store db.Store
}

// NewGlobalHandler creates a new handler that helps to manage all the handler functions
// and give db access to them
func NewGlobalHandler(store db.Store) *Handler {
	return &Handler{
		store: store,
	}
}

package handler

import (
	"github.com/Jayant-issar/severance-backend/internal/service"
)

//Returns a a new handler and store dependecies of all the handler function

type Handler struct {
	service *service.Service
}

// NewGlobalHandler creates a new handler that helps to manage all the handler functions
// and give db access to them
func NewGlobalHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

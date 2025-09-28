package service

import "github.com/Jayant-issar/severance-backend/internal/database/db"

type AssignmentService struct {
	store db.Store
}

func NewAssignmentService(store db.Store) *AssignmentService {
	return &AssignmentService{
		store: store,
	}
}

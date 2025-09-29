package service

import "github.com/Jayant-issar/severance-backend/internal/database/db"

type Service struct {
	User       *UserService
	Assignment *AssignmentService
}

func NewService(store db.Store) *Service {
	return &Service{
		User:       NewUserService(store),
		Assignment: NewAssignmentService(store),
	}
}

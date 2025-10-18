package service

import (
	"context"

	"github.com/Jayant-issar/severance-backend/severance-api/internal/database/db"
)

type UserService struct {
	store db.Store
}

func NewUserService(store db.Store) *UserService {
	return &UserService{
		store: store,
	}
}

func (us *UserService) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return us.store.CreateUser(ctx, arg)
}

func (us *UserService) GetUser(ctx context.Context, username string) (db.User, error) {
	return us.store.GetUser(ctx, username)
}

func (us *UserService) GetUserByID(ctx context.Context, id string) (db.User, error) {
	return us.store.GetUserByID(ctx, id)
}

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return us.store.GetUserByEmail(ctx, email)
}

func (us *UserService) ListUsers(ctx context.Context) ([]db.User, error) {
	return us.store.ListUsers(ctx)
}

func (us *UserService) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return us.store.UpdateUser(ctx, arg)
}

func (us *UserService) DeleteUser(ctx context.Context, id string) error {
	return us.store.DeleteUser(ctx, id)
}

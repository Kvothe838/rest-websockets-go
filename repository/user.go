package repository

import (
	"context"
	"rest-websockets-go/model"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	Close() error
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *model.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*model.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func Close() error {
	return implementation.Close()
}

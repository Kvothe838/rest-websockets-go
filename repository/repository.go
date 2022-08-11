package repository

import (
	"context"
	"rest-websockets-go/model"
)

type Repository interface {
	InsertUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	InsertPost(ctx context.Context, post *model.Post) error
	GetPostById(ctx context.Context, id string) (*model.Post, error)
	DeletePost(ctx context.Context, id string, userId string) error
	UpdatePost(ctx context.Context, post *model.Post, userId string) error
	//ListPost(ctx context.Context, userId string) ([]*models.Post, error)
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *model.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*model.User, error) {
	return implementation.GetUserByID(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertPost(ctx context.Context, post *model.Post) error {
	return implementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*model.Post, error) {
	return implementation.GetPostById(ctx, id)
}

func DeletePost(ctx context.Context, id string, userId string) error {
	return implementation.DeletePost(ctx, id, userId)
}

func UpdatePost(ctx context.Context, post *model.Post, userId string) error {
	return implementation.UpdatePost(ctx, post, userId)
}

// func ListPost(ctx context.Context, userId string) ([]*models.Post, error) {
// 	return implementation.ListPost(ctx, userId)
// }

func Close() error {
	return implementation.Close()
}

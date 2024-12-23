package domain

import "context"

type Repository interface {
	Insert(ctx context.Context, user *User) error
	FindOne(ctx context.Context, id string) (*User, error)
	FindOneByEmail(ctx context.Context, email string) (*User, error)
	Find(ctx context.Context, opts FindOptions) ([]*User, error)
}

type FindOptions struct {
	Email string
}

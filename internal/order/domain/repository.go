package domain

import "context"

type Repository interface {
	Insert(ctx context.Context, order *Order) error
	FindOne(ctx context.Context, id string) (*Order, error)
	Update(ctx context.Context, order *Order) error
	Find(ctx context.Context, opts *FindOptions) ([]*Order, error)
}

type FindOptions struct {
	Size      int
	UpdatedAt int64
}

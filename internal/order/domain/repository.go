package domain

import "context"

type Repository interface {
	Insert(ctx context.Context, order *Order) error
	FindOne(ctx context.Context, id string) (*Order, error)
}

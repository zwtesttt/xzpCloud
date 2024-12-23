package domain

import "context"

type Repository interface {
	FindOne(ctx context.Context, id string) (*Product, error)
	Insert(ctx context.Context, order *Product) error
}

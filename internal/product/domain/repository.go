package domain

import "context"

type Repository interface {
	FindOne(ctx context.Context, id string) (*Product, error)
	Insert(ctx context.Context, order *Product) error
	Find(ctx context.Context, opt *FindOptions) ([]*Product, error)
}

type FindOptions struct {
	UpdatedAt int64
	Size      int64
	//别的条件
}

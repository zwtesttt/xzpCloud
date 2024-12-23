package domain

import "context"

type Repository interface {
	Insert(ctx context.Context, vm *Vm) error
	FindOne(ctx context.Context, id string) (*Vm, error)
}

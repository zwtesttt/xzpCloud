package domain

import "context"

type Repository interface {
	Insert(ctx context.Context, vm *Vm) error
	FindOne(ctx context.Context, id string) (*Vm, error)
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context, opts *VmFindOptions) ([]*Vm, error)
}

type VmFindOptions struct {
	UserId string
	Name   string
}

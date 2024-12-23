package domain

import "context"

type Repository interface {
	FindOne(ctx context.Context, t RoleType) (*Role, error)
}

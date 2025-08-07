package domain

import (
	"context"

	"github.com/zwtesttt/xzpCloud/pkg/role"
)

type Repository interface {
	FindOne(ctx context.Context, t role.Type) (*Role, error)
}

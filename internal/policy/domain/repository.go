package domain

import "context"

type Repository interface {
	FindOne(ctx context.Context, sid string) (*ActionUrl, error)
}

package app

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/order/domain"
)

type GetOrderListInput struct {
	Size      int
	UpdatedAt int64
}

type GetOrderListHandler struct {
	orderRepo domain.Repository
}

func NewGetOrderListHandler(orderRepo domain.Repository) *GetOrderListHandler {
	return &GetOrderListHandler{
		orderRepo: orderRepo,
	}
}

func (h *GetOrderListHandler) Handle(ctx context.Context, input *GetOrderListInput) ([]*domain.Order, error) {
	return h.orderRepo.Find(ctx, &domain.FindOptions{
		Size:      input.Size,
		UpdatedAt: input.UpdatedAt,
	})
}

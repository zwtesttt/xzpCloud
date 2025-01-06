package app

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/order/domain"
)

type GetOrderInput struct {
	OrderId string
}

type GetOrderHandler struct {
	OrderRepository domain.Repository
}

func NewGetOrderHandler(orderRepository domain.Repository) *GetOrderHandler {
	return &GetOrderHandler{
		OrderRepository: orderRepository,
	}
}

func (h *GetOrderHandler) Handle(ctx context.Context, input *GetOrderInput) (*domain.Order, error) {
	return h.OrderRepository.FindOne(ctx, input.OrderId)
}

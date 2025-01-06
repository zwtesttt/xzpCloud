package app

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/order/domain"
	"github.com/zwtesttt/xzpCloud/pkg/api"
)

type CancelOrderInput struct {
	OrderId string
}

type CancelOrderHandler struct {
	orderRepo domain.Repository
}

func NewCancelOrderHandler(orderRepo domain.Repository) *CancelOrderHandler {
	return &CancelOrderHandler{
		orderRepo: orderRepo,
	}
}

func (h *CancelOrderHandler) Handle(ctx context.Context, input *CancelOrderInput) error {
	order, err := h.orderRepo.FindOne(ctx, input.OrderId)
	if err != nil {
		return err
	}

	if order.Status() != domain.OrderStatusPending {
		return api.FindCodeError(api.OrderStatusInvalid)
	}

	order.SetStatus(domain.OrderStatusCancelled)
	return h.orderRepo.Update(ctx, order)
}

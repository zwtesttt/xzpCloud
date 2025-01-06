package app

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/order/domain"
	"github.com/zwtesttt/xzpCloud/pkg/api"
)

type PayOrderInput struct {
	OrderId string
}

type PayOrderHandler struct {
	orderRepo domain.Repository
}

func NewPayOrderHandler(orderRepo domain.Repository) *PayOrderHandler {
	return &PayOrderHandler{
		orderRepo: orderRepo,
	}
}

func (h *PayOrderHandler) Handle(ctx context.Context, input *PayOrderInput) error {
	//TODO 集成第三方支付
	order, err := h.orderRepo.FindOne(ctx, input.OrderId)
	if err != nil {
		return err
	}

	if order.Status() != domain.OrderStatusPending {
		return api.FindCodeError(api.OrderStatusInvalid)
	}

	order.SetStatus(domain.OrderStatusPaid)
	return h.orderRepo.Update(ctx, order)
}

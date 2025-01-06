package app

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/order/domain"
)

type CreateOrderInput struct {
	UserId   string
	Products []*ProductItem
}

type ProductItem struct {
	ProductId string
	Quantity  int
	Price     float64
}

type CreateOrderHandler struct {
	orderRepo domain.Repository
}

func NewCreateOrderHandler(orderRepo domain.Repository) *CreateOrderHandler {
	return &CreateOrderHandler{
		orderRepo: orderRepo,
	}
}

func (h *CreateOrderHandler) Handle(ctx context.Context, input *CreateOrderInput) error {
	//TODO 校验商品合法
	//订单总金额
	total := 0.0
	//订单商品
	items := make([]*domain.Item, 0)
	for _, item := range input.Products {
		items = append(items, domain.NewOrderItem(item.ProductId, item.Quantity, item.Price))
		total += item.Price * float64(item.Quantity)
	}
	err := h.orderRepo.Insert(ctx, domain.NewOrder("", input.UserId, total, domain.OrderStatusPending, items, 0, 0, 0))
	if err != nil {
		return err
	}
	return nil
}

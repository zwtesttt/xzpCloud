package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/order/app"
	"github.com/zwtesttt/xzpCloud/pkg/api"
)

type GetOrderResponse struct {
	Id          string         `json:"id"`
	Products    []*ProductItem `json:"products"`
	TotalAmount float64        `json:"total_amount"`
	Status      int            `json:"status"`
	CreatedAt   int64          `json:"created_at"`
	UpdatedAt   int64          `json:"updated_at"`
}

func (h *Handler) GetOrder(ctx *gin.Context) {
	orderId := ctx.Param("id")
	if orderId == "" {
		api.RenderBadRequest(ctx)
		return
	}

	result, err := h.getOrderHandler.Handle(ctx, &app.GetOrderInput{OrderId: orderId})
	if err != nil {
		api.RenderError(ctx, err)
		return
	}

	items := make([]*ProductItem, 0)
	for _, item := range result.Items() {
		items = append(items, &ProductItem{ProductId: item.ProductId(), Quantity: item.Quantity(), Price: item.Price()})
	}

	resp := &GetOrderResponse{
		Id:          result.Id(),
		Products:    items,
		TotalAmount: result.TotalAmount(),
		Status:      int(result.Status()),
		CreatedAt:   result.CreatedAt(),
		UpdatedAt:   result.UpdatedAt(),
	}

	api.RenderSuccess(ctx, resp)
}

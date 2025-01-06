package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/order/app"
	"github.com/zwtesttt/xzpCloud/pkg/api"
	"strconv"
)

//type GetOrderListRequest struct {
//	UpdatedAt int `json:"last_updated_at"`
//	Size      int `json:"size"`
//}

type GetOrderListResponse struct {
	Orders []*OrderItem `json:"orders"`
}

type OrderItem struct {
	Id          string  `json:"id"`
	TotalAmount float64 `json:"total_amount"`
	Status      int     `json:"status"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}

func (h *Handler) GetOrderList(ctx *gin.Context) {
	updateTime := ctx.Query("last_updated_at")
	if updateTime == "" {
		api.RenderBadRequest(ctx)
		return
	}

	updatedTime, err := strconv.Atoi(updateTime)
	if err != nil {
		api.RenderBadRequest(ctx)
		return
	}

	size := ctx.Query("size")
	if size == "" {
		size = "20"
	}

	covSize, err := strconv.Atoi(updateTime)
	if err != nil {
		api.RenderBadRequest(ctx)
		return
	}

	result, err := h.getOrderListHandler.Handle(ctx, &app.GetOrderListInput{
		UpdatedAt: int64(updatedTime),
		Size:      covSize,
	})

	items := make([]*OrderItem, 0)
	for _, order := range result {
		items = append(items, &OrderItem{
			Id:          order.Id(),
			TotalAmount: order.TotalAmount(),
			Status:      int(order.Status()),
			CreatedAt:   order.CreatedAt(),
			UpdatedAt:   order.UpdatedAt(),
		})
	}

	api.RenderSuccess(ctx, GetOrderListResponse{
		Orders: items,
	})
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/order/app"
	"github.com/zwtesttt/xzpCloud/pkg/api"
)

type CreateOrderRequest struct {
	Products []*ProductItem `json:"products"`
	UserId   string         `json:"user_id"`
}

type ProductItem struct {
	Price     float64 `json:"price"`
	ProductId string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
}

func (h *Handler) CreateOrder(ctx *gin.Context) {
	//TODO 鉴权中间件
	//userId, ok := ctx.Get("user_id")
	//if !ok {
	//	api.RenderUnauthorized(ctx)
	//	return
	//}

	var req CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.RenderBadRequest(ctx)
		return
	}

	items := make([]*app.ProductItem, 0)
	for _, product := range req.Products {
		items = append(items, &app.ProductItem{ProductId: product.ProductId, Quantity: product.Quantity, Price: product.Price})
	}

	err := h.createOrderHandler.Handle(ctx, &app.CreateOrderInput{
		Products: items,
		UserId:   req.UserId,
	})
	if err != nil {
		api.RenderError(ctx, err)
		return
	}

	api.RenderSuccessNoBody(ctx)
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/order/app"
	"github.com/zwtesttt/xzpCloud/pkg/api"
)

func (h *Handler) PayOrder(ctx *gin.Context) {
	orderId := ctx.Param("id")
	if orderId == "" {
		api.RenderBadRequest(ctx)
		return
	}

	if err := h.payOrderHandler.Handle(ctx, &app.PayOrderInput{OrderId: orderId}); err != nil {
		api.RenderError(ctx, err)
		return
	}
	api.RenderSuccess(ctx, nil)
}

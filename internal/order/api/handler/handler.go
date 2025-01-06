package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/order/adapters"
	"github.com/zwtesttt/xzpCloud/internal/order/app"
	"github.com/zwtesttt/xzpCloud/pkg/db"
)

type Handler struct {
	*gin.Engine
	createOrderHandler  *app.CreateOrderHandler
	getOrderHandler     *app.GetOrderHandler
	cancelOrderHandler  *app.CancelOrderHandler
	payOrderHandler     *app.PayOrderHandler
	getOrderListHandler *app.GetOrderListHandler
}

func New() *Handler {
	orderRepo := adapters.NewOrderRepository(db.GetDB())

	h := &Handler{
		Engine:              gin.New(),
		createOrderHandler:  app.NewCreateOrderHandler(orderRepo),
		getOrderHandler:     app.NewGetOrderHandler(orderRepo),
		cancelOrderHandler:  app.NewCancelOrderHandler(orderRepo),
		payOrderHandler:     app.NewPayOrderHandler(orderRepo),
		getOrderListHandler: app.NewGetOrderListHandler(orderRepo),
	}

	o := h.Group("order")
	//查询订单列表
	o.GET("/", h.GetOrderList)
	//创建订单
	o.POST("/", h.CreateOrder)
	//查询订单详情
	o.GET("/:id", h.GetOrder)
	//取消订单
	o.POST("/:id/cancel", h.CancelOrder)
	//支付订单
	o.POST("/:id/pay", h.PayOrder)

	return h
}

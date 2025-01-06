package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/product/adapters"
	"github.com/zwtesttt/xzpCloud/internal/product/app"
	"github.com/zwtesttt/xzpCloud/pkg/db"
)

type Handler struct {
	*gin.Engine
	GetProductsHandler *app.GetProductsHandler
}

func New() *Handler {
	productRepo := adapters.NewProductRepository(db.GetDB())
	h := &Handler{
		Engine:             gin.New(),
		GetProductsHandler: app.NewGetProductsHandler(productRepo),
	}

	p := h.Group("/product")
	p.GET("/", h.GetProducts)
	return h
}

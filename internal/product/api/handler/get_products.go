package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/product/app"
	"github.com/zwtesttt/xzpCloud/pkg/api"
	"strconv"
)

type GetProductsRequest struct {
	UpdatedAt int `json:"last_updated_at"` // 上次请求的最后一条数据的创建时间，第一次请求为 0
	Size      int `json:"size"`            // 分页大小 默认20
}

type GetProductsResponse struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}

func (h *Handler) GetProducts(ctx *gin.Context) {
	timeStr := ctx.Query("last_updated_at")
	sizeStr := ctx.Query("size")

	if timeStr == "" {
		timeStr = "0"
	}
	updatedTime, err := strconv.Atoi(timeStr)
	if err != nil {
		api.RenderBadRequest(ctx)
		return
	}

	if sizeStr == "" {
		sizeStr = "20"
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		api.RenderBadRequest(ctx)
		return
	}

	req := GetProductsRequest{
		UpdatedAt: updatedTime,
		Size:      size,
	}

	result, err := h.GetProductsHandler.Handle(ctx, &app.GetProductsInput{
		UpdatedAt: int64(req.UpdatedAt),
		Size:      int64(req.Size),
	})
	if err != nil {
		api.RenderError(ctx, err)
		return
	}

	var resp GetProductsResponse
	for _, product := range result.Products {
		resp.Products = append(resp.Products, Product{
			ID:          product.Id(),
			Name:        product.Name(),
			Description: product.Description(),
			Price:       product.Price(),
			Stock:       product.Stock(),
			CreatedAt:   product.CreatedAt(),
			UpdatedAt:   product.UpdatedAt(),
		})
	}

	api.RenderSuccess(ctx, resp)
}

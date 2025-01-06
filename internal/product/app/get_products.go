package app

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/product/domain"
)

type GetProductsHandler struct {
	ProductRepository domain.Repository
}

func NewGetProductsHandler(productRepository domain.Repository) *GetProductsHandler {
	return &GetProductsHandler{
		ProductRepository: productRepository,
	}
}

type GetProductsInput struct {
	UpdatedAt int64 //游标
	Size      int64
}

type GetProductsOutput struct {
	Products []*domain.Product
}

func (h *GetProductsHandler) Handle(ctx context.Context, input *GetProductsInput) (*GetProductsOutput, error) {
	//插入数据
	//err := h.ProductRepository.Insert(ctx, domain.NewProduct("", fmt.Sprintf("虚拟机%s", time.Now().Format("20060102 15:04:05")), "虚拟机描述", 9.9, 10, 0, 0, 0))
	//if err != nil {
	//	return nil, err
	//}

	products, err := h.ProductRepository.Find(ctx, &domain.FindOptions{
		UpdatedAt: input.UpdatedAt,
		Size:      input.Size,
	})
	if err != nil {
		return nil, err
	}
	return &GetProductsOutput{
		Products: products,
	}, nil
}

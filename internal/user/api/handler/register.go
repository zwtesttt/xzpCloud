package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/user/app"
	"github.com/zwtesttt/xzpCloud/pkg/api"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

// TODO api接口文档注解
func (h *Handler) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.RenderBadRequest(ctx)
		return
	}

	if err := h.RegisterHandler.Handle(ctx, &app.RegisterInput{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}); err != nil {
		api.RenderError(ctx, err)
		return
	}

	api.RenderSuccess(ctx, nil)
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/vm/app"
	"github.com/zwtesttt/xzpCloud/pkg/api"
)

type CreateVmRequest struct {
	Name   string         `json:"name"`
	UserId string         `json:"user_id"`
	Config CreateVmConfig `json:"config"`
}

type CreateVmConfig struct {
	Cpu    int    `json:"cpu" binding:"required"`
	Memory string `json:"memory" binding:"required"`
	Disk   string `json:"disk" binding:"required"`
}

type CreateVmResponse struct {
	Id string `json:"id"`
}

func (h *Handler) CreateVm(ctx *gin.Context) {
	//TODO 鉴权中间件
	//userId, ok := ctx.Get("user_id")
	//if !ok {
	//	api.RenderUnauthorized(ctx)
	//	return
	//}
	var req CreateVmRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.RenderBadRequest(ctx)
		return
	}

	id, err := h.createVmHandler.Handle(ctx, &app.CreateVmInput{
		Name:   req.Name,
		UserId: req.UserId,
		Config: &app.CreateVmConfig{
			Cpu:    req.Config.Cpu,
			Memory: req.Config.Memory,
			Disk:   req.Config.Disk,
		},
	})
	if err != nil {
		api.RenderError(ctx, err)
		return
	}

	api.RenderSuccess(ctx, CreateVmResponse{Id: id})
}

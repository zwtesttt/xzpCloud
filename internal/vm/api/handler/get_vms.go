package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zwtesttt/xzpCloud/internal/vm/app"
	"github.com/zwtesttt/xzpCloud/pkg/api"
)

type GetVmsResponse struct {
	Total int   `json:"total"`
	Data  []*Vm `json:"data"`
}

type GetVmsRequest struct {
	UserId string // 用户ID，从token中拿
	//分页参数
	//筛选条件
}

type Vm struct {
	Config       *VmConfig `json:"config"`
	ExpirationAt int64     `json:"expiration_at"`
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Status       int       `json:"status"`
	Ip           string    `json:"ip"`
}

type VmConfig struct {
	Cpu    int    `json:"cpu"`
	Disk   string `json:"disk"`
	Memory string `json:"memory"`
}

func (h *Handler) GetVms(ctx *gin.Context) {
	//TODO UserId
	var req GetVmsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		api.RenderBadRequest(ctx)
		return
	}

	resp, err := h.getVmsHandler.Handle(ctx.Request.Context(), &app.GetVmsInput{UserId: ""})
	if err != nil {
		api.RenderError(ctx, err)
		return
	}

	var vms []*Vm
	for _, v := range resp.Vms {
		vms = append(vms, &Vm{
			Id:           v.Id,
			Name:         v.Name,
			Status:       v.Status,
			ExpirationAt: v.ExpirationAt,
			Ip:           v.Ip,
			Config: &VmConfig{
				Cpu:    v.Config.Cpu,
				Memory: v.Config.Memory,
				Disk:   v.Config.Disk,
			},
		})
	}

	api.RenderSuccess(ctx, &GetVmsResponse{
		Total: resp.Total,
		Data:  vms,
	})
}

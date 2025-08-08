package app

import (
	"context"

	"github.com/zwtesttt/xzpCloud/internal/vm/domain"
)

type GetVmsInput struct {
	UserId string
	//分页参数(偏移或者游标)
}

type GetVmsOutput struct {
	Vms   []*Vm
	Total int
}

type Vm struct {
	Id     string
	Name   string
	Status int
	Ip     string
	Config *VmConfig
	//过期时间
	ExpirationAt int64
}

type VmConfig struct {
	Cpu    int
	Memory string
	Disk   string
}

type GetVmsHandler struct {
	vmRepo domain.Repository
}

func NewGetVmsHandler(vmRepo domain.Repository) *GetVmsHandler {
	return &GetVmsHandler{vmRepo: vmRepo}
}

func (h *GetVmsHandler) Handle(ctx context.Context, input *GetVmsInput) (*GetVmsOutput, error) {
	vms, err := h.vmRepo.Find(ctx, &domain.VmFindOptions{UserId: input.UserId})
	if err != nil {
		return nil, err
	}

	var outputVms []*Vm
	for _, v := range vms {
		outputVms = append(outputVms, &Vm{
			Id:           v.Id(),
			Name:         v.Name(),
			Status:       int(v.Status()),
			Ip:           v.Ip(),
			ExpirationAt: v.ExpirationAt(),
			Config: &VmConfig{
				Cpu:    v.Config().Cpu(),
				Memory: v.Config().Memory(),
				Disk:   v.Config().Disk(),
			},
		})
	}

	return &GetVmsOutput{
		Vms:   outputVms,
		Total: len(outputVms),
	}, nil
}

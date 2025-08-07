package app

import (
	"context"

	"github.com/zwtesttt/xzpCloud/internal/vm/domain"
	"github.com/zwtesttt/xzpCloud/pkg/vmi"
)

type StopVmInput struct {
	Id     string
	UserId string
}

type StopVmHandler struct {
	vmRepo domain.Repository
	vmiCli vmi.VirtualMachineInterface
}

func NewStopVmHandler(vmRepo domain.Repository, vmiCli vmi.VirtualMachineInterface) *StopVmHandler {
	return &StopVmHandler{
		vmRepo: vmRepo,
		vmiCli: vmiCli,
	}
}

func (h *StopVmHandler) Handle(ctx context.Context, input *StopVmInput) error {
	// TODO 校验用户id
	vm, err := h.vmRepo.FindOne(ctx, input.Id)
	if err != nil {
		return err
	}

	err = h.vmiCli.Stop(ctx, &vmi.Config{
		Name:      vm.Name(),
		Namespace: "default",
	})
	if err != nil {
		return err
	}

	// 更新数据库中虚拟机状态为停止
	vm.SetStatus(domain.VmStatusStop)
	err = h.vmRepo.Update(ctx, vm)
	if err != nil {
		return err
	}

	return nil
}

package app

import (
	"context"

	"github.com/zwtesttt/xzpCloud/internal/vm/domain"
	"github.com/zwtesttt/xzpCloud/pkg/vmi"
)

type StartVmInput struct {
	Id     string
	UserId string
}

type StartVmHandler struct {
	vmRepo domain.Repository
	vmiCli vmi.VirtualMachineInterface
}

func NewStartVmHandler(vmRepo domain.Repository, vmiCli vmi.VirtualMachineInterface) *StartVmHandler {
	return &StartVmHandler{
		vmRepo: vmRepo,
		vmiCli: vmiCli,
	}
}

func (h *StartVmHandler) Handle(ctx context.Context, input *StartVmInput) error {
	// TODO 校验用户id
	vm, err := h.vmRepo.FindOne(ctx, input.Id)
	if err != nil {
		return err
	}

	err = h.vmiCli.Start(ctx, &vmi.Config{
		Name:      vm.Name(),
		Namespace: "default",
	})
	if err != nil {
		return err
	}

	// 更新数据库中虚拟机状态为运行中
	vm.SetStatus(domain.VmStatusStart)
	err = h.vmRepo.Update(ctx, vm)
	if err != nil {
		return err
	}

	return nil
}

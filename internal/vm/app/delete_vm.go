package app

import (
	"context"
	"github.com/zwtesttt/xzpCloud/internal/vm/domain"
	"github.com/zwtesttt/xzpCloud/pkg/vmi"
)

type DeleteVmInput struct {
	Id     string
	UserId string
}

type DeleteVmHandler struct {
	vmRepo domain.Repository
	vmiCli vmi.VirtualMachineInterface
}

func NewDeleteVmHandler(vmRepo domain.Repository, vmiCli vmi.VirtualMachineInterface) *DeleteVmHandler {
	return &DeleteVmHandler{
		vmRepo: vmRepo,
		vmiCli: vmiCli,
	}
}

func (h *DeleteVmHandler) Handle(ctx context.Context, input *DeleteVmInput) error {
	//TODO 校验用户id
	vm, err := h.vmRepo.FindOne(ctx, input.Id)
	if err != nil {
		return err
	}

	err = h.vmRepo.Delete(ctx, input.Id)
	if err != nil {
		return err
	}

	err = h.vmiCli.Delete(ctx, &vmi.Config{
		Name:      vm.Name(),
		Namespace: "default",
	})
	if err != nil {
		return err
	}
	return nil
}

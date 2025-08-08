package app

import (
	"context"
	"fmt"

	"github.com/zwtesttt/xzpCloud/internal/vm/domain"
	"github.com/zwtesttt/xzpCloud/pkg/vmi"
	v1 "kubevirt.io/api/core/v1"
)

type CreateVmInput struct {
	Name   string
	UserId string
	Config *CreateVmConfig
}

type CreateVmConfig struct {
	Cpu    int
	Memory string
	Disk   string
}

type CreateVmHandler struct {
	vmRepo domain.Repository
	vmiCli vmi.VirtualMachineInterface
}

func NewCreateVmHandler(vmRepo domain.Repository, vc vmi.VirtualMachineInterface) *CreateVmHandler {
	return &CreateVmHandler{vmRepo: vmRepo, vmiCli: vc}
}

func (h *CreateVmHandler) Handle(ctx context.Context, input *CreateVmInput) (string, error) {
	vm, err := h.vmiCli.Create(ctx, &vmi.Config{
		Name:      input.Name,
		Image:     "centos",
		Cpu:       input.Config.Cpu,
		Memory:    input.Config.Memory,
		Disk:      input.Config.Disk,
		Namespace: "default",
	})
	if err != nil {
		return "", err
	}

	vmi := vm.(*v1.VirtualMachine)

	fmt.Println(vmi)

	id, err := h.vmRepo.Insert(ctx, domain.NewVm("", input.Name, domain.VmStatusStart, input.UserId, "", domain.NewVmConfig(input.Config.Cpu, input.Config.Disk, input.Config.Memory), 0, 0, 0, 0))
	if err != nil {
		return "", err
	}
	return id, nil
}

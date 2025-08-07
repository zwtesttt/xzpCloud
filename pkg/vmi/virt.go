package vmi

import (
	"context"
	"fmt"
	"log"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/kubecli"

	"github.com/zwtesttt/xzpCloud/pkg/config"
)

type VirtHandler struct {
	kubeClient clientcmd.ClientConfig
	virtClient kubecli.KubevirtClient
}

var _ VirtualMachineInterface = &VirtHandler{}

func NewVirtHandler(cfg *config.Config) VirtualMachineInterface {
	//client := kubecli.DefaultClientConfig(&pflag.FlagSet{})
	//loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: cfg.KubeConfig.Path}
	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: cfg.KubeConfig.Path}

	configOverrides := &clientcmd.ConfigOverrides{}
	client := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	virtClient, err := kubecli.GetKubevirtClientFromClientConfig(client)
	if err != nil {
		log.Fatalf("cannot obtain KubeVirt client: %v\n", err)
	}
	return &VirtHandler{
		kubeClient: client,
		virtClient: virtClient,
	}
}

func (v *VirtHandler) Create(ctx context.Context, cfg *Config) (any, error) {
	//资源列表
	resources := make(apiv1.ResourceList)
	q, err := resource.ParseQuantity("64M")
	if err != nil {
		log.Fatalf("error in resource quantity: %v\n", err)
	}
	resources["memory"] = q

	//接口设置(接口是虚拟机的网卡)
	intf := make([]v1.Interface, 0)
	intf = append(intf, v1.Interface{
		Name: "default",
		InterfaceBindingMethod: v1.InterfaceBindingMethod{
			Masquerade: &v1.InterfaceMasquerade{},
		},
	})

	//网络设置
	netws := make([]v1.Network, 0)
	netws = append(netws, v1.Network{
		Name: "default",
		NetworkSource: v1.NetworkSource{Pod: &v1.PodNetwork{
			VMNetworkCIDR: "10.0.2.0/24",
		}},
	})
	create, err := v.virtClient.VirtualMachine(cfg.Namespace).Create(ctx, &v1.VirtualMachine{
		ObjectMeta: k8smetav1.ObjectMeta{Name: cfg.Name},
		Spec: v1.VirtualMachineSpec{
			Template: &v1.VirtualMachineInstanceTemplateSpec{
				ObjectMeta: k8smetav1.ObjectMeta{Labels: map[string]string{"kubevirt.io/size": "small", "kubevirt.io/domain": "testvm"}},
				Spec: v1.VirtualMachineInstanceSpec{
					Volumes: []v1.Volume{
						{
							Name: "containerdisk",
							VolumeSource: v1.VolumeSource{ContainerDisk: &v1.ContainerDiskSource{
								ImagePullPolicy: "IfNotPresent",
								Image:           "quay.io/kubevirt/cirros-container-disk-demo"},
							},
						},
						{
							Name: "cloudinitdisk",
							VolumeSource: v1.VolumeSource{
								CloudInitNoCloud: &v1.CloudInitNoCloudSource{UserDataBase64: "SGkuXG4="},
							},
						},
					},
					Networks: netws,
					Domain: v1.DomainSpec{
						Resources: v1.ResourceRequirements{
							Requests: resources,
						},
						Devices: v1.Devices{
							Interfaces: intf,
							Disks: []v1.Disk{
								{
									Name:       "containerdisk",
									DiskDevice: v1.DiskDevice{Disk: &v1.DiskTarget{Bus: "virtio"}},
								},
								{
									Name:       "cloudinitdisk",
									DiskDevice: v1.DiskDevice{Disk: &v1.DiskTarget{Bus: "virtio"}},
								},
							}},
					},
				},
			},
			RunStrategy: RunStrategyPtr(v1.RunStrategyAlways),
		}}, k8smetav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("error in create vm: %v", err)
	}
	return create, nil
}

func (v *VirtHandler) Delete(ctx context.Context, cfg *Config) error {
	err := v.virtClient.VirtualMachine(cfg.Namespace).Delete(ctx, cfg.Name, k8smetav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func BoolPtr(i bool) *bool { return &i }

func RunStrategyPtr(rs v1.VirtualMachineRunStrategy) *v1.VirtualMachineRunStrategy { return &rs }

func (v *VirtHandler) Start(ctx context.Context, cfg *Config) error {
	// 启动虚拟机：更新RunStrategy为Always
	vm, err := v.virtClient.VirtualMachine(cfg.Namespace).Get(ctx, cfg.Name, k8smetav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("error getting vm %s: %v", cfg.Name, err)
	}

	vm.Spec.RunStrategy = RunStrategyPtr(v1.RunStrategyAlways)
	_, err = v.virtClient.VirtualMachine(cfg.Namespace).Update(ctx, vm, k8smetav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("error starting vm %s: %v", cfg.Name, err)
	}

	fmt.Printf("Started VM %s\n", cfg.Name)
	return nil
}

func (v *VirtHandler) Stop(ctx context.Context, cfg *Config) error {
	// 停止虚拟机：更新RunStrategy为Halted
	vm, err := v.virtClient.VirtualMachine(cfg.Namespace).Get(ctx, cfg.Name, k8smetav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("error getting vm %s: %v", cfg.Name, err)
	}

	vm.Spec.RunStrategy = RunStrategyPtr(v1.RunStrategyHalted)
	_, err = v.virtClient.VirtualMachine(cfg.Namespace).Update(ctx, vm, k8smetav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("error stopping vm %s: %v", cfg.Name, err)
	}

	fmt.Printf("Stopped VM %s\n", cfg.Name)
	return nil
}

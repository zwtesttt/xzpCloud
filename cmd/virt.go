package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/spf13/pflag"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/kubecli"
)

func main() {

	// kubecli.DefaultClientConfig() prepares config using kubeconfig.
	// typically, you need to set env variable, KUBECONFIG=<path-to-kubeconfig>/.kubeconfig
	clientConfig := kubecli.DefaultClientConfig(&pflag.FlagSet{})

	// retrive default namespace.
	namespace, _, err := clientConfig.Namespace()
	if err != nil {
		log.Fatalf("error in namespace : %v\n", err)
	}

	// get the kubevirt client, using which kubevirt resources can be managed.
	virtClient, err := kubecli.GetKubevirtClientFromClientConfig(clientConfig)
	if err != nil {
		log.Fatalf("cannot obtain KubeVirt client: %v\n", err)
	}

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

	create, err := virtClient.VirtualMachine(namespace).Create(context.Background(), &v1.VirtualMachine{
		ObjectMeta: k8smetav1.ObjectMeta{Name: "testvm2"},
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
		log.Fatalf("error in create vm: %v\n", err)
		return
	}
	fmt.Printf("Created VM %s\n", create.ObjectMeta.Name)

	// Fetch list of VMs & VMIs
	vmList, err := virtClient.VirtualMachine(namespace).List(context.Background(), k8smetav1.ListOptions{})
	if err != nil {
		log.Fatalf("cannot obtain KubeVirt vm list: %v\n", err)
	}
	vmiList, err := virtClient.VirtualMachineInstance(namespace).List(context.Background(), k8smetav1.ListOptions{})
	if err != nil {
		log.Fatalf("cannot obtain KubeVirt vmi list: %v\n", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	fmt.Fprintln(w, "Type\tName\tNamespace\tStatus")

	for _, vm := range vmList.Items {
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\n", vm.Kind, vm.Name, vm.Namespace, vm.Status.Ready)
	}
	for _, vmi := range vmiList.Items {
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\n", vmi.Kind, vmi.Name, vmi.Namespace, vmi.Status.Phase)
	}
	w.Flush()
}

func BoolPtr(i bool) *bool { return &i }

func RunStrategyPtr(rs v1.VirtualMachineRunStrategy) *v1.VirtualMachineRunStrategy { return &rs }

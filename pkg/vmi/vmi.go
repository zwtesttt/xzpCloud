package vmi

import "context"

type VirtualMachineInterface interface {
	Create(ctx context.Context, cfg *Config) (interface{}, error)
	Delete(ctx context.Context, cfg *Config) error
}

type Config struct {
	Name      string
	Namespace string
	Image     string
	Cpu       int
	Memory    string
	Disk      string
}

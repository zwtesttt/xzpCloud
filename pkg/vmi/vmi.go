package vmi

import "context"

type VirtualMachineInterface interface {
	Create(ctx context.Context, cfg *Config) (any, error)
	Delete(ctx context.Context, cfg *Config) error
	Start(ctx context.Context, cfg *Config) error
	Stop(ctx context.Context, cfg *Config) error
}

type Config struct {
	Name      string
	Namespace string
	Image     string
	Cpu       int
	Memory    string
	Disk      string
}

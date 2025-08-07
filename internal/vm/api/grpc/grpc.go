package grpc

import (
	"github.com/zwtesttt/xzpCloud/pkg/config"
	"google.golang.org/grpc"
)

type VmGrpcServer struct {
}

func New(cfg *config.Config) *grpc.Server {
	s := grpc.NewServer()
	//vm := &VmGrpcServer{}
	//gegrpc.RegisterGreeterServer(s, vm)

	return s
}
